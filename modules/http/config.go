package http

type Config struct {
	Debug bool `mapstructure:"debug"`

	HTTP struct {
		Bind      string `mapstructure:"bind" default:"0.0.0.0:8080"`
		KeepAlive bool   `mapstructure:"keep_alive" default:"true"`
	} `mapstructure:"http"`
}
