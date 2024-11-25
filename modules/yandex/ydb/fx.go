package ydb

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/config"
)

const ModuleName = "ydb"

var Module = fx.Module(
	ModuleName,

	config.Inject(new(Config)),

	fx.Provide(
		NewDriver,
		NewGormDriver,
	),

	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),
)
