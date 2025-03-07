package pprof

// Config
type Config struct {
	Monitoring struct {
		Pprof struct {
			Enabled bool `mapstructure:"enabled"`
		} `mapstructure:"pprof"`
	} `mapstructure:"monitoring"`
}
