package logger

type Config struct {
	Debug bool `mapstructure:"debug"`

	Logger struct {
		Level string `mapstructure:"level" validate:"required" default:"info"`
	} `mapstructure:"logger"`
}
