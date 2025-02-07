package fasthttp

import "time"

// Config
type Config struct {
	FastHTTP struct {
		Client struct {
			// Host fasthttp client
			Host struct {
				MaxConnsPerHost     int           `mapstructure:"max_conns_per_host" default:"1024"`
				MaxConnDuration     time.Duration `mapstructure:"max_conn_duration" default:"30s"`
				MaxConnWaitTimeout  time.Duration `mapstructure:"max_conn_wait_timeout" default:"10ms"`
				MaxIdleConnDuration time.Duration `mapstructure:"max_idle_conn_durartion" default:"10s"`
				ReadTimeout         time.Duration `mapstructure:"read_timeout" default:"1s"`
				WriteTimeout        time.Duration `mapstructure:"write_timeout" default:"1s"`
				ReadBufferSize      int           `mapstructure:"read_buffer_size" default:"1024"`
				WriteBufferSize     int           `mapstructure:"write_buffer_size" default:"1024"`
			} `mapstructure:"host"`

			// Pipeline fasthttp client
			Pipeline struct {
				MaxConns            int           `mapstructure:"max_conns" default:"1024"`
				MaxPendingRequests  int           `mapstructure:"max_pending_requests" default:"1024"`
				MaxBatchDelay       time.Duration `mapstructure:"max_batch_delay" default:"0"`
				MaxIdleConnDuration time.Duration `mapstructure:"max_idle_conn_durartion" default:"10s"`
				ReadTimeout         time.Duration `mapstructure:"read_timeout" default:"1s"`
				WriteTimeout        time.Duration `mapstructure:"write_timeout" default:"1s"`
				ReadBufferSize      int           `mapstructure:"read_buffer_size" default:"1024"`
				WriteBufferSize     int           `mapstructure:"write_buffer_size" default:"1024"`
			} `mapstructure:"pipeline"`
		} `mapstructure:"client"`
	} `mapstructure:"fasthttp"`
}
