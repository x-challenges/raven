package graphql

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/config"
)

const ModuleName = "graphql"

var Module = fx.Module(
	ModuleName,

	config.Inject(new(Config)),

	// public usage
	fx.Provide(
		NewServer,
	),

	// private usage
	fx.Provide(
		fx.Private,

		newRecoveryFn,
		newErrorPresenterFn,
		newAPQCacheAdapter,
	),

	// register http controllers
	fx.Invoke(
		NewQueryController,
		NewPlaygroundController,
	),

	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),
)
