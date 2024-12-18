package tracing

// Config
type Config struct {
	Service struct {
		Name string `mapstructure:"name" default:"app"`
	} `mapstructure:"service"`

	Monitoring struct {
		Tracing struct {
			Enabled  bool   `mapstructure:"enabled"`
			Endpoint string `mapstructure:"endpoint" default:"http://0.0.0.0:4317"`
		} `mapstructure:"tracing"`
	} `mapstructure:"monitoring"`
}
