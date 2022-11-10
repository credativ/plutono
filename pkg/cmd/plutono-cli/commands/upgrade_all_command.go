package commands

import (
	"github.com/credativ/plutono/pkg/cmd/plutono-cli/logger"
	"github.com/credativ/plutono/pkg/cmd/plutono-cli/models"
	"github.com/credativ/plutono/pkg/cmd/plutono-cli/services"
	"github.com/credativ/plutono/pkg/cmd/plutono-cli/utils"
	"github.com/hashicorp/go-version"
)

func shouldUpgrade(installed string, remote *models.Plugin) bool {
	installedVersion, err := version.NewVersion(installed)
	if err != nil {
		return false
	}

	latest := latestSupportedVersion(remote)
	latestVersion, err := version.NewVersion(latest.Version)
	if err != nil {
		return false
	}
	return installedVersion.LessThan(latestVersion)
}

func (cmd Command) upgradeAllCommand(c utils.CommandLine) error {
	pluginsDir := c.PluginDirectory()

	localPlugins := services.GetLocalPlugins(pluginsDir)

	remotePlugins, err := cmd.Client.ListAllPlugins(c.String("repo"))
	if err != nil {
		return err
	}

	pluginsToUpgrade := make([]models.InstalledPlugin, 0)

	for _, localPlugin := range localPlugins {
		for _, p := range remotePlugins.Plugins {
			remotePlugin := p
			if localPlugin.ID != remotePlugin.ID {
				continue
			}
			if shouldUpgrade(localPlugin.Info.Version, &remotePlugin) {
				pluginsToUpgrade = append(pluginsToUpgrade, localPlugin)
			}
		}
	}

	for _, p := range pluginsToUpgrade {
		logger.Infof("Updating %v \n", p.ID)

		err := services.RemoveInstalledPlugin(pluginsDir, p.ID)
		if err != nil {
			return err
		}

		err = InstallPlugin(p.ID, "", c, cmd.Client)
		if err != nil {
			return err
		}
	}

	return nil
}
