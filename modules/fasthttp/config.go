package fasthttp

import "time"

// Config
type Config struct {
	FastHTTP struct {
		Client struct {
			MaxConnsPerHost    int           `mapstructure:"max_conns_per_host" default:"1024"`
			MaxConnDuration    time.Duration `mapstructure:"max_conn_duration" default:"30s"`
			MaxConnWaitTimeout time.Duration `mapstructure:"max_conn_wait_timeout" default:"10ms"`
			ReadTimeout        time.Duration `mapstructure:"read_timeout" default:"1s"`
			WriteTimeout       time.Duration `mapstructure:"write_timeout" default:"1s"`
		} `mapstructure:"client"`
	} `mapstructure:"fasthttp"`
}
