package trigger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const ModuleName = "trigger"

var Module = fx.Module(
	ModuleName,

	fx.Provide(
		NewFactory,
	),

	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),
)
