package logger

import (
	"os"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type FxLoggerParam struct {
	fx.In

	Config *Config
	Cores  []zapcore.Core `group:"zapcores"`
}

// NewLogger construct.
func NewLogger(p FxLoggerParam) (*zap.Logger, error) {
	var (
		level zapcore.Level
		atom  = zap.NewAtomicLevel()
		core  zapcore.Core
		err   error

		opts = []zap.Option{
			zap.AddCaller(),
			zap.AddStacktrace(zapcore.ErrorLevel),
		}
	)

	// set config level
	if level, err = zapcore.ParseLevel(p.Config.Logger.Level); err != nil {
		return nil, err
	}

	// set debug if provided
	if p.Config.Debug {
		level = zapcore.DebugLevel
	}

	atom.SetLevel(level)

	switch p.Config.Debug {
	case true:
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(
				zap.NewDevelopmentEncoderConfig(),
			),
			zapcore.Lock(os.Stdout),
			atom,
		)

	default:
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(
				zap.NewProductionEncoderConfig(),
			),
			zapcore.Lock(os.Stdout),
			atom,
		)
	}

	return zap.New(
		zapcore.NewTee(append(p.Cores, core)...), opts...), nil
}
