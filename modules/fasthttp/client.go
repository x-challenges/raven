package fasthttp

import (
	"github.com/valyala/fasthttp"
)

// NewClient
func NewClient(config *Config) *fasthttp.Client {
	return &fasthttp.Client{
		MaxConnsPerHost:          config.FastHTTP.Client.MaxConnsPerHost,
		MaxConnDuration:          config.FastHTTP.Client.MaxConnDuration,
		MaxConnWaitTimeout:       config.FastHTTP.Client.MaxConnWaitTimeout,
		ReadTimeout:              config.FastHTTP.Client.ReadTimeout,
		WriteTimeout:             config.FastHTTP.Client.WriteTimeout,
		NoDefaultUserAgentHeader: true,
	}
}
