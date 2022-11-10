package commands

import (
	"errors"

	"github.com/credativ/plutono/pkg/cmd/plutono-cli/logger"
	"github.com/credativ/plutono/pkg/cmd/plutono-cli/utils"
)

func validateVersionInput(c utils.CommandLine) error {
	arg := c.Args().First()
	if arg == "" {
		return errors.New("please specify plugin to list versions for")
	}

	return nil
}

func (cmd Command) listVersionsCommand(c utils.CommandLine) error {
	if err := validateVersionInput(c); err != nil {
		return err
	}

	pluginToList := c.Args().First()

	plugin, err := cmd.Client.GetPlugin(pluginToList, c.String("repo"))
	if err != nil {
		return err
	}

	for _, i := range plugin.Versions {
		logger.Infof("%v\n", i.Version)
	}

	return nil
}
