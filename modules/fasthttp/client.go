package fasthttp

import (
	"github.com/valyala/fasthttp"
)

// Client
type Client = fasthttp.Client

// NewClient
func NewClient(config *Config) *Client {
	var cfg = config.FastHTTP.Client.Host

	return &fasthttp.Client{
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
	}
}

// PipelineClient
type PipelineClient = fasthttp.PipelineClient

// NewPipelineClient
func NewPipelineClient(logger *LoggerAdapter, config *Config) *PipelineClient {
	var cfg = config.FastHTTP.Client.Pipeline

	return &fasthttp.PipelineClient{
		Logger:                        logger,
		MaxConns:                      cfg.MaxConns,
		MaxPendingRequests:            cfg.MaxPendingRequests,
		MaxBatchDelay:                 cfg.MaxBatchDelay,
		MaxIdleConnDuration:           cfg.MaxIdleConnDuration,
		ReadBufferSize:                cfg.ReadBufferSize,
		WriteBufferSize:               cfg.WriteBufferSize,
		ReadTimeout:                   cfg.ReadTimeout,
		WriteTimeout:                  cfg.WriteTimeout,
		DialDualStack:                 true,
		NoDefaultUserAgentHeader:      true,
		DisableHeaderNamesNormalizing: true,
		DisablePathNormalizing:        true,
	}
}
