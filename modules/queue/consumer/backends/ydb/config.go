package ydb

type Config struct {
	Consumer string `mapstructure:"consumer" validate:"required"`
	Topic    string `mapstructure:"topic" validate:"required"`
}
