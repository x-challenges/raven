package fasthttp

import "github.com/valyala/fasthttp"

var (
	AcquireRequest = fasthttp.AcquireRequest
	ReleaseRequest = fasthttp.ReleaseRequest

	AcquireResponse = fasthttp.AcquireResponse
	ReleaseResponse = fasthttp.ReleaseResponse
)

type (
	Request  = fasthttp.Request
	Response = fasthttp.Response
)
