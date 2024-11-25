package gorm

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const ModuleName = "gorm"

var Module = fx.Module(
	ModuleName,

	fx.Provide(
		NewDriver,
	),

	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),
)
