package flood

import "time"

type Config struct {
	Flood struct{} `mapstructure:"flood"`
}

type Flood struct {
	Worker struct {
		Queue    string        `mapstructure:"queue" default:"default"`
		Timeout  time.Duration `mapstructure:"timeout" default:"30s"`
		MaxRetry int           `mapstructure:"max_retry"`
	} `mapstructure:"worker"`

	Cron struct {
		Spec string `mapstructure:"spec" validate:"omitempty,cron" default:""`
	} `mapstructure:"cron"`
}
