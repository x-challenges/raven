package graphql

import (
	"time"

	"github.com/x-challenges/raven/modules/cache"
)

type Config struct {
	Graphql struct {
		Tracing         bool `mapstructure:"tracing" default:"false"`
		Intorspection   bool `mapstructure:"introspection" default:"false"`
		ComplexityLimit int  `mapstructure:"complexity_limit" default:"500"`

		Query struct {
			Path string `mapstructure:"path" validate:"required" default:"/graphql"`
		} `mapstructure:"query"`

		Playground struct {
			Enabled   bool   `mapstructure:"enabled"`
			Path      string `mapstructure:"path" validate:"required" default:"/playground"`
			QueryPath string `mapstructure:"query_path" validate:"required" default:"/graphql"`
		} `mapstructure:"playground"`

		Websocket struct {
			Ping      time.Duration `mapstructure:"ping" default:"5s"`
			KeepAlive time.Duration `mapstructure:"keep_alive" default:"15s"`
		} `mapstructure:"websocket"`

		APQ struct {
			Enabled bool `mapstructure:"enabled" default:"false"`
			Limit   int  `mapstructure:"limit" default:"1000"`
		} `mapstructure:"apq"`

		Errors struct {
			Debug bool `mapstructure:"debug" default:"false"`
		} `mapstructure:"errors"`

		Caches struct {
			APQ cache.Options `mapstructure:"apq"`
		} `mapstructure:"caches"`
	} `mapstructure:"graphql"`
}
