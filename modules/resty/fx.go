package resty

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/config"
)

const ModuleName = "resty"

var Module = fx.Module(
	ModuleName,

	config.Inject(new(Config)),

	fx.Provide(
		NewLoggerAdapter,
		NewClient,
	),

	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),
)
