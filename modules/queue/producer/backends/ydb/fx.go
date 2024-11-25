package ydb

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const ModuleName = "ydb"

var Module = fx.Module(
	ModuleName,

	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),
)
