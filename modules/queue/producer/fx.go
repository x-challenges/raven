package producer

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/yandex/ydb"
)

const ModuleName = "producer"

var Module = fx.Module(
	ModuleName,

	ydb.Module,

	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),
)
