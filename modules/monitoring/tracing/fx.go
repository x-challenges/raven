package tracing

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/config"
)

// ModuleName
const ModuleName = "tracer"

// Module
var Module = fx.Module(
	ModuleName,

	config.Inject(new(Config)),

	// force
	fx.Invoke(InvokeTraceProvider),

	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),
)
