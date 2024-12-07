package fasthttp

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ModuleName
const ModuleName = "fasthttp"

// Module
var Module = fx.Module(
	ModuleName,

	// public usage
	fx.Provide(
		NewClient,
	),

	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),
)
