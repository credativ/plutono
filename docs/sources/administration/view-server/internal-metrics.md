+++
title = "Internal Plutono metrics"
description = "Internal metrics exposed by Plutono"
keywords = ["plutono", "metrics", "internal metrics"]
aliases = ["/docs/plutono/latest/admin/metrics/"]
weight = 200
+++

# Internal Plutono metrics

Plutono collects some metrics about itself internally. Plutono supports pushing metrics to Graphite or exposing them to be scraped by Prometheus.

For more information about configuration options related to Plutono metrics, refer to [metrics]({{< relref "../../administration/configuration/#metrics" >}}) and [metrics.graphite]({{< relref "../../administration/configuration/#metrics-graphite" >}}) in [Configuration]({{< relref "../../administration/configuration.md" >}}).

## Available metrics

When enabled, Plutono exposes a number of metrics, including:

- Active Plutono instances
- Number of dashboards, users, and playlists
- HTTP status codes
- Requests by routing group
- Plutono active alerts
- Plutono performance

## Pull metrics from Plutono into Prometheus

These instructions assume you have already added Prometheus as a data source in Plutono.

1. Enable Prometheus to scrape metrics from Plutono. In your configuration file (`plutono.ini` or `custom.ini` depending on your operating system) remove the semicolon to enable the following configuration options:

   ```
   # Metrics available at HTTP API Url /metrics
   [metrics]
   # Disable / Enable internal metrics
   enabled           = true

   # Disable total stats (stat_totals_*) metrics to be generated
   disable_total_stats = false
   ```

1. (optional) If you want to require authorization to view the metrics endpoint, then uncomment and set the following options:

   ```
   basic_auth_username =
   basic_auth_password =
   ```

1. Restart Plutono. Plutono now exposes metrics at http://localhost:3000/metrics.
1. Add the job to your prometheus.yml file.
   Example:

   ```
   - job_name: 'plutono_metrics'

      scrape_interval: 15s
      scrape_timeout: 5s

      static_configs:
        - targets: ['localhost:3000']
   ```
1. Restart Prometheus. Your new job should appear on the Targets tab.
1. In Plutono, hover your mouse over the **Configuration** (gear) icon on the left sidebar and then click **Data Sources**.
1. Select the **Prometheus** data source.
1. On the Dashboards tab, **Import** the Plutono metrics dashboard. All scraped Plutono metrics are available in the dashboard.

## View Plutono metrics in Graphite

These instructions assume you have already added Graphite as a data source in Plutono.

1. Enable sending metrics to Graphite. In your configuration file (`plutono.ini` or `custom.ini` depending on your operating system) remove the semicolon to enable the following configuration options:

   ```
   # Metrics available at HTTP API Url /metrics
   [metrics]
   # Disable / Enable internal metrics
   enabled           = true

   # Disable total stats (stat_totals_*) metrics to be generated
   disable_total_stats = false
   ```

1. Enable [metrics.graphite] options:
   ```
   # Send internal metrics to Graphite
   [metrics.graphite]
   # Enable by setting the address setting (ex localhost:2003)
   address = <hostname or ip>:<port#>
   prefix = prod.plutono.%(instance_name)s.
   ```

1. Restart Plutono. Plutono now exposes metrics at http://localhost:3000/metrics and sends them to the Graphite location you specified.
