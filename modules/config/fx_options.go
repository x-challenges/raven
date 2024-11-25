package config

type options struct {
	optionalPrefix string
}

type Option func(*options)

// WithOptionalPrefix add optional prefix for configuration.
func WithOptionalPrefix(prefix string) Option {
	return func(o *options) {
		o.optionalPrefix = prefix
	}
}
