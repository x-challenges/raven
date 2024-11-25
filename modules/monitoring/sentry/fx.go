package sentry

import (
	"github.com/getsentry/sentry-go"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/config"
)

var logger *zap.Logger

const ModuleName = "sentry"

var Module = fx.Module(
	ModuleName,

	config.Inject(new(Config)),

	// provide new sentry client
	fx.Provide(
		fx.Annotate(
			NewClient,

			// graceful shutdown
			fx.OnStop(
				func(config *Config, client *sentry.Client) {
					_ = client.Flush(config.Sentry.Flush)
				},
			),
		),
	),

	// provide new zapcore with sentry error logger to DI graph
	fx.Provide(
		fx.Annotate(
			NewZapcore,
			fx.ResultTags(`group:"zapcores"`),
		),
	),

	fx.Populate(&logger), // extract logger to module global

	// force
	fx.Invoke(func(_ *sentry.Client) {}),

	fx.ErrorHook(
		NewFxErrorHandler(),
	),
)
