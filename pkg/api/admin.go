package api

import (
	"github.com/credativ/plutono/pkg/api/response"
	"github.com/credativ/plutono/pkg/bus"
	"github.com/credativ/plutono/pkg/models"
	"github.com/credativ/plutono/pkg/setting"
)

func AdminGetSettings(c *models.ReqContext) response.Response {
	settings := make(map[string]interface{})

	for _, section := range setting.Raw.Sections() {
		jsonSec := make(map[string]interface{})
		settings[section.Name()] = jsonSec

		for _, key := range section.Keys() {
			keyName := key.Name()
			value := setting.RedactedValue(setting.EnvKey(section.Name(), key.Name()), key.Value())

			jsonSec[keyName] = value
		}
	}

	return response.JSON(200, settings)
}

func AdminGetStats(c *models.ReqContext) response.Response {
	statsQuery := models.GetAdminStatsQuery{}

	if err := bus.Dispatch(&statsQuery); err != nil {
		return response.Error(500, "Failed to get admin stats from database", err)
	}

	return response.JSON(200, statsQuery.Result)
}
