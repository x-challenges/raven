package limiter

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const ModuleName = "limiter"

// Module provided to fx.
var Module = fx.Module(
	ModuleName,

	// public uasge
	fx.Provide(
		New,
	),

	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),
)
