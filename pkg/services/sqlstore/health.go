package sqlstore

import (
	"github.com/credativ/plutono/pkg/bus"
	"github.com/credativ/plutono/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetDBHealthQuery)
}

func GetDBHealthQuery(query *models.GetDBHealthQuery) error {
	_, err := x.Exec("SELECT 1")
	return err
}
