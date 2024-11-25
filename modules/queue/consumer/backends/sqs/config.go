package sqs

type Config struct {
	URL string `mapstructure:"url" validate:"required"`
}
