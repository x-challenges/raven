package cache

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/cache/stores"
)

const ModuleName = "cache"

var Module = fx.Module(
	ModuleName,

	stores.Module,

	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),
)
