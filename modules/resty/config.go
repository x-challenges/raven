package resty

import "time"

type Config struct {
	Debug bool `mapstructure:"debug" default:"false"`

	Resty struct {
		Trace   bool          `mapstructure:"trace" default:"false"`
		Timeout time.Duration `mapstructure:"timeout" default:"15s"`

		Retry struct {
			Count int `mapstructure:"count" default:"0"`
		} `mapstructure:"retry"`
	} `mapstructure:"resty"`
}
