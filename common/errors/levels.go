package errors

import (
	"go.uber.org/zap/zapcore"
)

type Level string

const (
	System Level = "system"
	Client Level = "client"
)

var (
	AsSystem = WithLevelOption(System)
	AsClient = WithLevelOption(Client)
)

func (l Level) IsSystem() bool { return l == System }

func (l Level) ZapLogLevel() zapcore.Level {
	switch l {
	case System:
		return zapcore.ErrorLevel
	case Client:
		return zapcore.WarnLevel
	default:
		return zapcore.ErrorLevel
	}
}

func GetLevel(err error) Level {
	var (
		level = System
	)

	for err != nil {
		if e, ok := err.(*Error); ok {
			if e.Level != nil {
				level = *e.Level
				break
			}
		}

		err = Unwrap(err)
	}
	return level
}
