package pprof

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/config"
)

// ModuleName
const ModuleName = "pprof"

// Module
var Module = fx.Module(
	ModuleName,

	// config
	config.Inject(new(Config)),

	// invoke
	fx.Invoke(NewDebugController),

	// decorate
	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),
)
