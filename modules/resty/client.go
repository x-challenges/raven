package resty

import (
	"github.com/go-resty/resty/v2"
)

func NewClient(logger resty.Logger, config *Config) *resty.Client {
	var (
		client = resty.New()
	)

	// debug mode
	client.SetDebug(config.Debug)

	// set zap logger adapter
	client.SetLogger(logger)

	// enable tracing
	if config.Resty.Trace {
		client.EnableTrace()
	}

	// set timeout duration
	if t := config.Resty.Timeout; t.Seconds() > 0 {
		client.SetTimeout(t)
	}

	// set retry count
	if c := config.Resty.Retry.Count; c > 0 {
		client.SetRetryCount(c)
	}

	return client
}
