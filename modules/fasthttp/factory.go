package fasthttp

import (
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
	"go.uber.org/zap"
)

// Factory
type Factory interface {
	// New
	New(opts ...FactoryOptionFunc) *Client
}

// Factory interface implementation
type factory struct {
	logger *zap.Logger
	config *Config
}

var _ Factory = (*factory)(nil)

// NewFactory
func NewFactory(logger *zap.Logger, config *Config) (*factory, error) {
	return &factory{
		logger: logger,
		config: config,
	}, nil
}

// New implements Factory interface
func (f *factory) New(opts ...FactoryOptionFunc) *Client {
	var (
		options = NewFactoryOptions().Apply(opts...)
		dialer  fasthttp.DialFunc
	)

	// set default config
	if options.Config == nil {
		options.Config = f.config
	}

	var cfg = options.Config.FastHTTP.Client.Host

	// init proxy dialer if endabled
	if proxy := options.Proxy; proxy != "" {
		dialer = fasthttpproxy.FasthttpHTTPDialerTimeout(proxy, cfg.ReadTimeout)
	}

	// new fast http client
	var client = &fasthttp.Client{
		MaxConnsPerHost:               cfg.MaxConnsPerHost,
		MaxConnDuration:               cfg.MaxConnDuration,
		MaxConnWaitTimeout:            cfg.MaxConnWaitTimeout,
		MaxIdleConnDuration:           cfg.MaxIdleConnDuration,
		ReadTimeout:                   cfg.ReadTimeout,
		WriteTimeout:                  cfg.WriteTimeout,
		ReadBufferSize:                cfg.ReadBufferSize,
		WriteBufferSize:               cfg.WriteBufferSize,
		NoDefaultUserAgentHeader:      true,
		DisableHeaderNamesNormalizing: true,
		DisablePathNormalizing:        true,
		DialDualStack:                 true,
		Dial:                          dialer,
	}

	return client
}

// FactoryOptions
type FactoryOptions struct {
	Config *Config `json:"config"`
	Proxy  string  `json:"proxy"`
}

// NewFactoryOptions
func NewFactoryOptions() *FactoryOptions {
	return &FactoryOptions{}
}

// FactoryOptionFunc
type FactoryOptionFunc func(*FactoryOptions)

// Apply
func (fo *FactoryOptions) Apply(opts ...FactoryOptionFunc) *FactoryOptions {
	for _, opt := range opts {
		opt(fo)
	}
	return fo
}

// WithConfig
func WithConfig(config *Config) FactoryOptionFunc {
	return func(fo *FactoryOptions) {
		fo.Config = config
	}
}

// WithProxy
func WithProxy(proxy string) FactoryOptionFunc {
	return func(fo *FactoryOptions) {
		fo.Proxy = proxy
	}
}
