package asynq

import "time"

type Config struct {
	Asynq struct {
		Server struct {
			Concurrency    int            `mapstructure:"concurrency" validate:"required" default:"1"`
			Queues         map[string]int `mapstructure:"queues"`
			StrictPriority bool           `mapstructure:"strict_priority" default:"true"`

			Shutdown struct {
				Timeout time.Duration `mapstructure:"timeout" default:"30s"`
			} `mapstructure:"shutdown"`
		} `mapstructure:"server"`

		Scheduler struct {
			Location string `mapstructure:"location" default:"UTC"`
		} `mapstructure:"scheduler"`
	} `mapstructure:"asynq"`
}
