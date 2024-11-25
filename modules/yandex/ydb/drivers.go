package ydb

import (
	"context"

	ydbgorm "github.com/ydb-platform/gorm-driver"
	environ "github.com/ydb-platform/ydb-go-sdk-auth-environ"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"gorm.io/gorm"
)

func NewDriver(config *Config) (*ydb.Driver, error) {
	return ydb.Open(
		context.Background(),
		config.YDB.DSN,
		environ.WithEnvironCredentials(),
	)
}

func NewGormDriver(config *Config) gorm.Dialector {
	return ydbgorm.Open(
		config.YDB.DSN,
		ydbgorm.WithTablePathPrefix(config.YDB.TablePathPrefix),
		ydbgorm.WithMaxOpenConns(config.YDB.MaxOpenConns),
		ydbgorm.WithMaxIdleConns(config.YDB.MaxIdleConns),
		ydbgorm.WithConnMaxIdleTime(config.YDB.MaxIdleTime),
		ydbgorm.With(environ.WithEnvironCredentials()),
	)
}
