package aws

type Config struct {
	AWS struct {
		Endpoint struct {
			URL    string `mapstructure:"url" validate:"required"`
			Region string `mapstructure:"region" validate:"required"`
		} `mapstructure:"endpoint"`
	} `mapstructure:"aws"`
}
