package backends

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/queue/consumer/backends/sqs"
	"github.com/x-challenges/raven/modules/queue/consumer/backends/ydb"
)

const ModuleName = "backends"

var Module = fx.Module(
	ModuleName,

	// backends
	sqs.Module,
	ydb.Module,

	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),
)
