package plugins

import (
	"context"
	"encoding/json"
	"path/filepath"

	"github.com/credativ/plutono/pkg/infra/log"
	"github.com/credativ/plutono/pkg/plugins/backendplugin"
	"github.com/credativ/plutono/pkg/plugins/backendplugin/grpcplugin"
	"github.com/credativ/plutono/pkg/plugins/backendplugin/pluginextensionv2"
	"github.com/credativ/plutono/pkg/util/errutil"
	pluginModel "github.com/grafana/grafana-plugin-model/go/renderer"
)

type RendererPlugin struct {
	FrontendPluginBase

	Executable           string `json:"executable,omitempty"`
	GrpcPluginV1         pluginModel.RendererPlugin
	GrpcPluginV2         pluginextensionv2.RendererPlugin
	backendPluginManager backendplugin.Manager
}

func (r *RendererPlugin) Load(decoder *json.Decoder, base *PluginBase, backendPluginManager backendplugin.Manager) error {
	if err := decoder.Decode(r); err != nil {
		return err
	}

	if err := r.registerPlugin(base); err != nil {
		return err
	}

	r.backendPluginManager = backendPluginManager

	cmd := ComposePluginStartCommand("plugin_start")
	fullpath := filepath.Join(r.PluginDir, cmd)
	factory := grpcplugin.NewRendererPlugin(r.Id, fullpath, grpcplugin.PluginStartFuncs{
		OnLegacyStart: r.onLegacyPluginStart,
		OnStart:       r.onPluginStart,
	})
	if err := backendPluginManager.Register(r.Id, factory); err != nil {
		return errutil.Wrapf(err, "Failed to register backend plugin")
	}

	Renderer = r
	return nil
}

func (r *RendererPlugin) Start(ctx context.Context) error {
	if err := r.backendPluginManager.StartPlugin(ctx, r.Id); err != nil {
		return errutil.Wrapf(err, "Failed to start renderer plugin")
	}

	return nil
}

func (r *RendererPlugin) onLegacyPluginStart(pluginID string, client *grpcplugin.LegacyClient, logger log.Logger) error {
	r.GrpcPluginV1 = client.RendererPlugin
	return nil
}

func (r *RendererPlugin) onPluginStart(pluginID string, client *grpcplugin.Client, logger log.Logger) error {
	r.GrpcPluginV2 = client.RendererPlugin
	return nil
}
