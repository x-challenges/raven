package dataloaders

const (
	limit = 30
)

// Options struct
type Options struct {
	// Limit batch fn
	Limit int
}

func NewOptions() *Options {
	return &Options{
		Limit: limit,
	}
}

// Option type
type Option func(*Options)

// WithLimit option setter
func WithLimit(limit int) Option {
	return func(o *Options) {
		o.Limit = limit
	}
}
