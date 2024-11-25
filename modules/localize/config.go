package localize

type Config struct {
	Langs []string `envconfig:"LOCALIZE_LANGS" required:"true" default:"EN,RU"`
	Path  string   `envconfig:"LOCALIZE_PATH" default:"./assets/localizations"`
}

func NewConfig() *Config {
	return &Config{}
}
