package fasthttp

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/config"
)

// ModuleName
const ModuleName = "fasthttp"

// Module
var Module = fx.Module(
	ModuleName,

	config.Inject(new(Config)),

	// public usage
	fx.Provide(
		// Client
		NewClient,

		// Factory
		fx.Annotate(NewFactory, fx.As(new(Factory))),
	),

	// private usage
	fx.Provide(
		fx.Private,

		// LoggerAdapter
		NewLoggerAdapter,
	),

	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),
)
