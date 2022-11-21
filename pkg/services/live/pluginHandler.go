package live

import (
	"github.com/centrifugal/centrifuge"
	"github.com/credativ/plutono/pkg/models"
	"github.com/credativ/plutono/pkg/plugins"
)

// PluginHandler manages all the `plutono/dashboard/*` channels
type PluginHandler struct {
	Plugin *plugins.PluginBase
}

// GetHandlerForPath called on init
func (h *PluginHandler) GetHandlerForPath(path string) (models.ChannelHandler, error) {
	return h, nil // all dashboards share the same handler
}

// OnSubscribe for now allows anyone to subscribe
func (h *PluginHandler) OnSubscribe(c *centrifuge.Client, e centrifuge.SubscribeEvent) (centrifuge.SubscribeReply, error) {
	return centrifuge.SubscribeReply{}, nil
}

// OnPublish checks if a message from the websocket can be broadcast on this channel
func (h *PluginHandler) OnPublish(c *centrifuge.Client, e centrifuge.PublishEvent) (centrifuge.PublishReply, error) {
	return centrifuge.PublishReply{}, nil // broadcast any event
}
