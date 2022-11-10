package plugins

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/credativ/plutono/pkg/infra/log"
	"github.com/credativ/plutono/pkg/setting"
	"github.com/hashicorp/go-version"
)

var (
	httpClient = http.Client{Timeout: 10 * time.Second}
)

type PlutonoNetPlugin struct {
	Slug    string `json:"slug"`
	Version string `json:"version"`
}

type GithubLatest struct {
	Stable  string `json:"stable"`
	Testing string `json:"testing"`
}

func getAllExternalPluginSlugs() string {
	var result []string
	for _, plug := range Plugins {
		if plug.IsCorePlugin {
			continue
		}

		result = append(result, plug.Id)
	}

	return strings.Join(result, ",")
}

func (pm *PluginManager) checkForUpdates() {
	if !setting.CheckForUpdates {
		return
	}

	pm.log.Debug("Checking for updates")

	pluginSlugs := getAllExternalPluginSlugs()
	resp, err := httpClient.Get("https://grafana.com/api/plugins/versioncheck?slugIn=" + pluginSlugs + "&plutonoVersion=" + setting.BuildVersion)
	if err != nil {
		log.Tracef("Failed to get plugins repo from grafana.com, %v", err.Error())
		return
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Warn("Failed to close response body", "err", err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Tracef("Update check failed, reading response from grafana.com, %v", err.Error())
		return
	}

	gNetPlugins := []PlutonoNetPlugin{}
	err = json.Unmarshal(body, &gNetPlugins)
	if err != nil {
		log.Tracef("Failed to unmarshal plugin repo, reading response from grafana.com, %v", err.Error())
		return
	}

	for _, plug := range Plugins {
		for _, gplug := range gNetPlugins {
			if gplug.Slug == plug.Id {
				plug.PlutonoNetVersion = gplug.Version

				plugVersion, err1 := version.NewVersion(plug.Info.Version)
				gplugVersion, err2 := version.NewVersion(gplug.Version)

				if err1 != nil || err2 != nil {
					plug.PlutonoNetHasUpdate = plug.Info.Version != plug.PlutonoNetVersion
				} else {
					plug.PlutonoNetHasUpdate = plugVersion.LessThan(gplugVersion)
				}
			}
		}
	}

	resp2, err := httpClient.Get("https://raw.githubusercontent.com/plutono/plutono/master/latest.json")
	if err != nil {
		log.Tracef("Failed to get latest.json repo from github.com: %v", err.Error())
		return
	}
	defer func() {
		if err := resp2.Body.Close(); err != nil {
			pm.log.Warn("Failed to close response body", "err", err)
		}
	}()
	body, err = ioutil.ReadAll(resp2.Body)
	if err != nil {
		log.Tracef("Update check failed, reading response from github.com, %v", err.Error())
		return
	}

	var githubLatest GithubLatest
	err = json.Unmarshal(body, &githubLatest)
	if err != nil {
		log.Tracef("Failed to unmarshal github.com latest, reading response from github.com: %v", err.Error())
		return
	}

	if strings.Contains(setting.BuildVersion, "-") {
		pm.PlutonoLatestVersion = githubLatest.Testing
		pm.PlutonoHasUpdate = !strings.HasPrefix(setting.BuildVersion, githubLatest.Testing)
	} else {
		pm.PlutonoLatestVersion = githubLatest.Stable
		pm.PlutonoHasUpdate = githubLatest.Stable != setting.BuildVersion
	}

	currVersion, err1 := version.NewVersion(setting.BuildVersion)
	latestVersion, err2 := version.NewVersion(pm.PlutonoLatestVersion)
	if err1 == nil && err2 == nil {
		pm.PlutonoHasUpdate = currVersion.LessThan(latestVersion)
	}
}
