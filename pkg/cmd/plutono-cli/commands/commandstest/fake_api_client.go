package commandstest

import (
	"os"

	"github.com/credativ/plutono/pkg/cmd/plutono-cli/models"
)

type FakePlutonoComClient struct {
	GetPluginFunc      func(pluginId, repoUrl string) (models.Plugin, error)
	DownloadFileFunc   func(pluginName string, tmpFile *os.File, url string, checksum string) (err error)
	ListAllPluginsFunc func(repoUrl string) (models.PluginRepo, error)
}

func (client *FakePlutonoComClient) GetPlugin(pluginID, repoUrl string) (models.Plugin, error) {
	if client.GetPluginFunc != nil {
		return client.GetPluginFunc(pluginID, repoUrl)
	}

	return models.Plugin{}, nil
}

func (client *FakePlutonoComClient) DownloadFile(pluginName string, tmpFile *os.File, url string, checksum string) (err error) {
	if client.DownloadFileFunc != nil {
		return client.DownloadFileFunc(pluginName, tmpFile, url, checksum)
	}

	return nil
}

func (client *FakePlutonoComClient) ListAllPlugins(repoURL string) (models.PluginRepo, error) {
	if client.ListAllPluginsFunc != nil {
		return client.ListAllPluginsFunc(repoURL)
	}
	return models.PluginRepo{}, nil
}
