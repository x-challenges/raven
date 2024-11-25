package sentry

import "time"

type Config struct {
	Debug bool `mapstructure:"debug"`

	Sentry struct {
		DSN                 string        `mapstructure:"dsn"`
		SampleRate          float64       `mapstructure:"sample_rate" default:"1.0"`
		Flush               time.Duration `mapstructure:"flush" default:"5s"`
		SkipSSLVerification bool          `mapstructure:"skip_ssl" default:"true"`

		Release struct {
			Service     string `mapstructure:"service" default:"gatefi"`
			Version     string `mapstructure:"version" default:"local"`
			Environment string `mapstructure:"environment" default:"local"`
		} `mapstructure:"release"`
	} `mapstructure:"sentry"`
}
