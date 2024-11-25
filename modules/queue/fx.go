package queue

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/queue/consumer"
)

const ModuleName = "queue"

var Module = fx.Module(
	ModuleName,

	consumer.Module,

	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),
)
