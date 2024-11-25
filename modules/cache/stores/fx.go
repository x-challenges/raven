package stores

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const ModuleName = "stores"

var Module = fx.Module(
	ModuleName,

	fx.Provide(
		NewFactory,
	),

	// store factories
	fx.Provide(
		fx.Private,

		fx.Annotate(newMemoryFactory, fx.ResultTags(`group:"cache:stores:factory"`)),
		fx.Annotate(newRedisFactory, fx.ResultTags(`group:"cache:stores:factory"`)),
	),

	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),
)
