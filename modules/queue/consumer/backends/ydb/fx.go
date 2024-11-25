package ydb

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/queue/consumer/backends/backend"
)

const ModuleName = "ydb"

var Module = fx.Module(
	ModuleName,

	fx.Provide(
		fx.Annotate(
			NewFactory,
			fx.As(new(backend.Factory)),
			fx.ResultTags(`group:"queue:consumer:backend:factory"`),
		),
	),

	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),
)
