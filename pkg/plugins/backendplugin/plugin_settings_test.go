package backendplugin

import (
	"sort"
	"testing"

	"github.com/credativ/plutono/pkg/setting"
	"github.com/stretchr/testify/require"
)

func TestPluginSettings(t *testing.T) {
	t.Run("Should only extract from sections beginning with 'plugin.' in config", func(t *testing.T) {
		cfg := &setting.Cfg{
			PluginSettings: setting.PluginSettings{
				"plugin": map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			},
		}

		ps := extractPluginSettings(cfg)
		require.Len(t, ps, 1)
		require.Len(t, ps["plugin"], 2)

		t.Run("Should skip path setting", func(t *testing.T) {
			cfg.PluginSettings["plugin"]["path"] = "value"
			ps := extractPluginSettings(cfg)
			require.Len(t, ps["plugin"], 2)
		})

		t.Run("Should skip id setting", func(t *testing.T) {
			cfg.PluginSettings["plugin"]["id"] = "value"
			ps := extractPluginSettings(cfg)
			require.Len(t, ps["plugin"], 2)
		})

		t.Run("Should return expected environment variables from plugin settings ", func(t *testing.T) {
			ps := extractPluginSettings(cfg)
			env := ps["plugin"].ToEnv("PL_PLUGIN", []string{"PL_VERSION=6.7.0"})
			sort.Strings(env)
			require.Len(t, env, 3)
			require.EqualValues(t, []string{"PL_PLUGIN_KEY1=value1", "PL_PLUGIN_KEY2=value2", "PL_VERSION=6.7.0"}, env)
		})
	})
}
