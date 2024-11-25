package redis

import "time"

type Config struct {
	Redis struct {
		Addresses []string `mapstructure:"addresses" default:"0.0.0.0:6379"`
		Database  int      `mapstructure:"database" default:"0"`

		// Client configuration
		Client struct {
			Name string `mapstructure:"name" default:"redis-fx"`
		} `mapstructure:"client"`

		// Dial policy
		Dial struct {
			Timeout time.Duration `mapstructure:"timeout"`
		} `mapstructure:"dial"`

		// Read policy
		Read struct {
			Timeout time.Duration `mapstructure:"timeout"`
		} `mapstructure:"read"`

		// Write policy
		Write struct {
			Timeout time.Duration `mapstructure:"timeout"`
		} `mapstructure:"write"`

		// Retry policy
		Retry struct {
			Count      int           `mapstructure:"count" default:"0"`
			MinBackoff time.Duration `mapstructure:"min_backoff"`
			MaxBackoff time.Duration `mapstructure:"max_backoff"`
		} `mapstructure:"retry"`

		// Pool configuration
		Pool struct {
			Size            int           `mapstructure:"size"`
			Timeout         time.Duration `mapstructure:"timeout"`
			MinIdleConns    int           `mapstructure:"min_idle_conns"`
			MaxIdleConns    int           `mapstructure:"max_idle_conns"`
			MaxActiveConns  int           `mapstructure:"max_active_conns"`
			ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
		} `mapstructure:"pool"`
	} `mapstructure:"redis"`
}
