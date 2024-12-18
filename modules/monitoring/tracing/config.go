package tracing

// Config
type Config struct {
	Service struct {
		Name string `mapstructure:"name" default:"app"`
	} `mapstructure:"service"`

	Monitoring struct {
		Tracing struct {
			Enabled bool `mapstructure:"enabled"`
		} `mapstructure:"tracing"`
	} `mapstructure:"monitoring"`
}
