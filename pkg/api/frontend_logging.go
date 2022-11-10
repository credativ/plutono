package api

import (
	"github.com/getsentry/sentry-go"
	"github.com/credativ/plutono/pkg/api/frontendlogging"
	"github.com/credativ/plutono/pkg/api/response"
	"github.com/credativ/plutono/pkg/infra/log"
	"github.com/credativ/plutono/pkg/models"
)

var frontendLogger = log.New("frontend")

type frontendLogMessageHandler func(c *models.ReqContext, event frontendlogging.FrontendSentryEvent) response.Response

func NewFrontendLogMessageHandler(store *frontendlogging.SourceMapStore) frontendLogMessageHandler {
	return func(c *models.ReqContext, event frontendlogging.FrontendSentryEvent) response.Response {
		var msg = "unknown"

		if len(event.Message) > 0 {
			msg = event.Message
		} else if event.Exception != nil && len(event.Exception.Values) > 0 {
			msg = event.Exception.Values[0].FmtMessage()
		}

		var ctx = event.ToLogContext(store)

		switch event.Level {
		case sentry.LevelError:
			frontendLogger.Error(msg, ctx)
		case sentry.LevelWarning:
			frontendLogger.Warn(msg, ctx)
		case sentry.LevelDebug:
			frontendLogger.Debug(msg, ctx)
		default:
			frontendLogger.Info(msg, ctx)
		}

		return response.Success("ok")
	}
}
