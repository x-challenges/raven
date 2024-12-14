package metrics

// Config
type Config struct {
	Monitoring struct {
		Metrics struct {
			ExporterURL string `mapstructure:"exporter_url" validate:"required" default:"/metrics"`
		} `mapstructure:"metrics"`
	} `mapstructure:"monitoring"`
}
