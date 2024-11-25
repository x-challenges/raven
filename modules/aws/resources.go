package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func NewAwsConfig(cfg *Config) (aws.Config, error) {
	return config.LoadDefaultConfig(
		context.Background(),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(func(_, _ string, _ ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:           cfg.AWS.Endpoint.URL,
					SigningRegion: cfg.AWS.Endpoint.Region,
				}, nil
			}),
		),
	)
}

func NewClient(cfg aws.Config) *sqs.Client {
	return sqs.NewFromConfig(cfg)
}
