package consumer

import (
	"github.com/x-challenges/raven/modules/queue/consumer/backends/backend"
)

type Config struct {
	Queue struct {
		Consumers map[string]struct {
			backend.Config `mapstructure:",squash"`
		} `mapstructure:"consumers" validate:"dive,required"`
	} `mapstructure:"queue"`
}
