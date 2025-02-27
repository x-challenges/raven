package fasthttp

import (
	"github.com/valyala/fasthttp"
)

// Client
type Client = fasthttp.Client

// NewClient
func NewClient(factory Factory) *Client {
	return factory.New()
}
