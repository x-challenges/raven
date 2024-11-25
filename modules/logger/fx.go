package logger

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/x-challenges/raven/modules/config"
)

const ModuleName = "logger"

// Module provided to fx.
var Module = fx.Module(
	ModuleName,

	config.Inject(new(Config)),

	fx.Provide(
		fx.Annotate(
			NewLogger,

			fx.OnStop(
				func(logger *zap.Logger) { _ = logger.Sync() },
			),
		),
	),

	fx.Invoke(func(_ *zap.Logger) {}),
)

var FxLogger = fx.WithLogger(
	func(logger *zap.Logger) fxevent.Logger {
		var (
			fxlogger = &fxevent.ZapLogger{Logger: logger}
		)

		fxlogger.UseLogLevel(zapcore.DebugLevel)

		return fxlogger
	},
)
