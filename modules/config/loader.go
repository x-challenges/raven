package config

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
)

var validate = validator.New()

// Config type.
type Config interface{}

// Loader interface.
type Loader interface {
	// Load configuration
	Load(Config) error
}

// Loader interface implementation.
type loader struct {
	prefix string
	viper  *viper.Viper
}

// NewLoader construct.
func NewLoader(prefix string) func(viper *viper.Viper) (Loader, error) {
	return func(viper *viper.Viper) (Loader, error) {
		loader := &loader{
			prefix: prefix,
			viper:  viper,
		}

		// setup files loader
		if len(Files) > 0 {
			if err := setupConfigFiles(loader.viper, Files...); err != nil {
				return nil, err
			}
		}

		// setup environment loader
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv()

		return loader, nil
	}
}

// Run load parsed configurations to config.
func (l *loader) Load(c Config) error {
	var (
		err error
	)

	// make env binder
	if err = bindenvs(l.viper, c); err != nil {
		return err
	}

	// unmarshal values from viper
	if err = l.unmarshal(c); err != nil {
		return err
	}

	// set defaults
	defaults.SetDefaults(c)

	// validate config
	return validate.Struct(c)
}

// Unmarshal configuration.
func (l *loader) unmarshal(c Config) error {
	var (
		err error
	)

	// try load config with prefix
	if l.prefix != "" {
		if err = l.viper.UnmarshalKey(l.prefix, c); err != nil {
			return err
		}
	}

	// try load config without prefix
	if err = l.viper.Unmarshal(c); err != nil {
		return err
	}

	return nil
}
