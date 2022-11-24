package usagestats

import (
	"context"
	"errors"
	"testing"

	"github.com/credativ/plutono/pkg/services/alerting"
	"github.com/credativ/plutono/pkg/services/licensing"
	"github.com/stretchr/testify/require"

	"github.com/credativ/plutono/pkg/bus"
	"github.com/credativ/plutono/pkg/models"
	"github.com/credativ/plutono/pkg/services/sqlstore"
	"github.com/credativ/plutono/pkg/setting"
	"github.com/stretchr/testify/assert"
)

// This is to ensure that the interface contract is held by the implementation
func Test_InterfaceContractValidity(t *testing.T) {
	newUsageStats := func() UsageStats {
		return &UsageStatsService{}
	}
	v, ok := newUsageStats().(*UsageStatsService)

	assert.NotNil(t, v)
	assert.True(t, ok)
}

func TestMetrics(t *testing.T) {
	t.Run("When updating total stats", func(t *testing.T) {
		uss := &UsageStatsService{
			Bus: bus.New(),
			Cfg: setting.NewCfg(),
		}
		uss.Cfg.MetricsEndpointEnabled = true
		uss.Cfg.MetricsEndpointDisableTotalStats = false
		getSystemStatsWasCalled := false
		uss.Bus.AddHandler(func(query *models.GetSystemStatsQuery) error {
			query.Result = &models.SystemStats{}
			getSystemStatsWasCalled = true
			return nil
		})

		t.Run("When metrics is disabled and total stats is enabled", func(t *testing.T) {
			uss.Cfg.MetricsEndpointEnabled = false
			uss.Cfg.MetricsEndpointDisableTotalStats = false
			t.Run("Should not update stats", func(t *testing.T) {
				uss.updateTotalStats()

				assert.False(t, getSystemStatsWasCalled)
			})
		})

		t.Run("When metrics is enabled and total stats is disabled", func(t *testing.T) {
			uss.Cfg.MetricsEndpointEnabled = true
			uss.Cfg.MetricsEndpointDisableTotalStats = true

			t.Run("Should not update stats", func(t *testing.T) {
				uss.updateTotalStats()

				assert.False(t, getSystemStatsWasCalled)
			})
		})

		t.Run("When metrics is disabled and total stats is disabled", func(t *testing.T) {
			uss.Cfg.MetricsEndpointEnabled = false
			uss.Cfg.MetricsEndpointDisableTotalStats = true

			t.Run("Should not update stats", func(t *testing.T) {
				uss.updateTotalStats()

				assert.False(t, getSystemStatsWasCalled)
			})
		})

		t.Run("When metrics is enabled and total stats is enabled", func(t *testing.T) {
			uss.Cfg.MetricsEndpointEnabled = true
			uss.Cfg.MetricsEndpointDisableTotalStats = false

			t.Run("Should update stats", func(t *testing.T) {
				uss.updateTotalStats()

				assert.True(t, getSystemStatsWasCalled)
			})
		})
	})

	t.Run("When registering a metric", func(t *testing.T) {
		uss := &UsageStatsService{
			Bus:             bus.New(),
			Cfg:             setting.NewCfg(),
			externalMetrics: make(map[string]MetricFunc),
		}
		metricName := "stats.test_metric.count"

		t.Run("Adds a new metric to the external metrics", func(t *testing.T) {
			uss.RegisterMetric(metricName, func() (interface{}, error) {
				return 1, nil
			})

			metric, _ := uss.externalMetrics[metricName]()
			assert.Equal(t, 1, metric)
		})

		t.Run("When metric already exists", func(t *testing.T) {
			uss.RegisterMetric(metricName, func() (interface{}, error) {
				return 1, nil
			})

			metric, _ := uss.externalMetrics[metricName]()
			assert.Equal(t, 1, metric)

			t.Run("Overrides the metric", func(t *testing.T) {
				uss.RegisterMetric(metricName, func() (interface{}, error) {
					return 2, nil
				})
				newMetric, _ := uss.externalMetrics[metricName]()
				assert.Equal(t, 2, newMetric)
			})
		})
	})

	t.Run("When getting usage report", func(t *testing.T) {
		uss := &UsageStatsService{
			Bus:                bus.New(),
			Cfg:                setting.NewCfg(),
			SQLStore:           sqlstore.InitTestDB(t),
			License:            &licensing.OSSLicensingService{},
			AlertingUsageStats: &alertingUsageMock{},
			externalMetrics:    make(map[string]MetricFunc),
		}
		metricName := "stats.test_metric.count"

		uss.Bus.AddHandler(func(query *models.GetSystemStatsQuery) error {
			query.Result = &models.SystemStats{}
			return nil
		})

		uss.Bus.AddHandler(func(query *models.GetDataSourceStatsQuery) error {
			query.Result = []*models.DataSourceStats{}
			return nil
		})

		uss.Bus.AddHandler(func(query *models.GetDataSourcesByTypeQuery) error {
			query.Result = []*models.DataSource{}
			return nil
		})

		uss.Bus.AddHandler(func(query *models.GetDataSourceAccessStatsQuery) error {
			query.Result = []*models.DataSourceAccessStats{}
			return nil
		})

		uss.Bus.AddHandler(func(query *models.GetAlertNotifierUsageStatsQuery) error {
			query.Result = []*models.NotifierUsageStats{}
			return nil
		})

		createConcurrentTokens(t, uss.SQLStore)

		t.Run("Should include metrics for concurrent users", func(t *testing.T) {
			report, err := uss.GetUsageReport(context.Background())
			require.NoError(t, err)

			assert.Equal(t, int32(1), report.Metrics["stats.auth_token_per_user_le_3"])
			assert.Equal(t, int32(2), report.Metrics["stats.auth_token_per_user_le_6"])
			assert.Equal(t, int32(3), report.Metrics["stats.auth_token_per_user_le_9"])
			assert.Equal(t, int32(4), report.Metrics["stats.auth_token_per_user_le_12"])
			assert.Equal(t, int32(5), report.Metrics["stats.auth_token_per_user_le_15"])
			assert.Equal(t, int32(6), report.Metrics["stats.auth_token_per_user_le_inf"])
		})

		t.Run("Should include external metrics", func(t *testing.T) {
			uss.RegisterMetric(metricName, func() (interface{}, error) {
				return 1, nil
			})

			report, err := uss.GetUsageReport(context.Background())
			assert.Nil(t, err, "Expected no error")

			metric := report.Metrics[metricName]
			assert.Equal(t, 1, metric)
		})
	})

	t.Run("When registering external metrics", func(t *testing.T) {
		uss := &UsageStatsService{
			Bus:             bus.New(),
			Cfg:             setting.NewCfg(),
			externalMetrics: make(map[string]MetricFunc),
		}
		metrics := map[string]interface{}{"stats.test_metric.count": 1, "stats.test_metric_second.count": 2}
		extMetricName := "stats.test_external_metric.count"

		t.Run("Should add to metrics", func(t *testing.T) {
			uss.RegisterMetric(extMetricName, func() (interface{}, error) {
				return 1, nil
			})

			uss.registerExternalMetrics(metrics)

			assert.Equal(t, 1, metrics[extMetricName])
		})

		t.Run("When loading a metric results to an error", func(t *testing.T) {
			uss.RegisterMetric(extMetricName, func() (interface{}, error) {
				return 1, nil
			})
			extErrorMetricName := "stats.test_external_metric_error.count"

			t.Run("Should not add it to metrics", func(t *testing.T) {
				uss.RegisterMetric(extErrorMetricName, func() (interface{}, error) {
					return 1, errors.New("some error")
				})

				uss.registerExternalMetrics(metrics)

				extErrorMetric := metrics[extErrorMetricName]
				extMetric := metrics[extMetricName]

				assert.Nil(t, extErrorMetric, "Invalid metric should not be added")
				assert.Equal(t, 1, extMetric)
				assert.Len(t, metrics, 3, "Expected only one available metric")
			})
		})
	})
}

type alertingUsageMock struct{}

func (aum *alertingUsageMock) QueryUsageStats() (*alerting.UsageStats, error) {
	return &alerting.UsageStats{
		DatasourceUsage: map[string]int{
			"prometheus":         1,
			"graphite":           2,
			"mysql":              5,
			"unknown-datasource": 90,
		},
	}, nil
}
