package backend

import "context"

type Type = string

type Config struct {
	// Client  string                 `mapstructure:"client" validate:"required" default:"todo"`
	Backend map[string]interface{} `mapstructure:"backend" validate:"required"`
}

type Callback func(ctx context.Context, messages ...interface{}) error

type Backend interface {
	// Run
	Run(ctx context.Context, callback Callback) error

	// Close
	Close(ctx context.Context) error
}
