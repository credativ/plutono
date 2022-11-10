package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/credativ/plutono/pkg/cmd/plutono-cli/services"
	"github.com/credativ/plutono/pkg/cmd/plutono-cli/utils"
)

var removePlugin func(pluginPath, id string) error = services.RemoveInstalledPlugin

func (cmd Command) removeCommand(c utils.CommandLine) error {
	pluginPath := c.PluginDirectory()

	plugin := c.Args().First()
	if plugin == "" {
		return errors.New("missing plugin parameter")
	}

	err := removePlugin(pluginPath, plugin)

	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return fmt.Errorf("plugin does not exist")
		}

		return err
	}

	return nil
}
