package sentry

import (
	"github.com/TheZeroSlave/zapsentry"
	"github.com/getsentry/sentry-go"
	"go.uber.org/zap/zapcore"
)

func NewZapcore(client *sentry.Client) (zapcore.Core, error) {
	return zapsentry.NewCore(
		zapsentry.Configuration{
			Level:             zapcore.ErrorLevel,
			EnableBreadcrumbs: true,
			BreadcrumbLevel:   zapcore.InfoLevel,
		}, zapsentry.NewSentryClientFromClient(client),
	)
}
