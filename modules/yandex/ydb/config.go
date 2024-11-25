package ydb

import "time"

type Config struct {
	YDB struct {
		DSN             string        `mapstructure:"dsn" validate:"required" default:"grpc://localhost:2136/local"`
		TablePathPrefix string        `mapstructure:"table_path_prefix"`
		MaxOpenConns    int           `mapstructure:"max_open_conns" default:"10"`
		MaxIdleConns    int           `mapstructure:"max_idle_conns" default:"2"`
		MaxIdleTime     time.Duration `mapstructure:"max_idle_time" default:"60s"`
	} `mapstructure:"ydb"`
}
