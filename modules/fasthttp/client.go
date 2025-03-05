package fasthttp

import (
	"github.com/valyala/fasthttp"
)

// Client
type Client = fasthttp.Client

// PipelineClient
type PipelineClient = fasthttp.PipelineClient

// NewClient
func NewClient(factory Factory) *Client {
	return factory.Client()
}
