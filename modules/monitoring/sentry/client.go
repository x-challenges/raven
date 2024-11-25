package sentry

import (
	"crypto/tls"
	"net/http"
	"os"

	"github.com/getsentry/sentry-go"
)

// NewClient construct.
func NewClient(config *Config) (*sentry.Client, error) {
	var (
		httpClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: config.Sentry.SkipSSLVerification, //nolint:gosec
				},
			},
		}

		opts = sentry.ClientOptions{
			Debug:            config.Debug,
			Dsn:              config.Sentry.DSN,
			SampleRate:       config.Sentry.SampleRate,
			Release:          config.Sentry.Release.Service,
			Dist:             config.Sentry.Release.Version,
			Environment:      config.Sentry.Release.Environment,
			AttachStacktrace: true,
			HTTPClient:       httpClient,
		}

		client *sentry.Client
		err    error
	)

	if serverName, err := os.Hostname(); err == nil {
		opts.ServerName = serverName
	}

	// if b := p.Build; b != nil {
	// 	opts.Release = b.Service
	// 	opts.Dist = b.Version
	// 	opts.Environment = b.Environment
	// }

	hub := sentry.CurrentHub()

	if client, err = sentry.NewClient(opts); err != nil {
		return nil, err
	}

	hub.BindClient(client)

	return client, nil
}
