package localize

// Options contains data for localizing
type Options struct {
	TemplateData map[string]interface{}
	PluralCount  int
}

type Option func(*Options)

func WithTemplateData(data map[string]interface{}) Option {
	return func(o *Options) {
		o.TemplateData = data
	}
}

func WithPluralCount(count int) Option {
	return func(o *Options) {
		o.PluralCount = count
	}
}
