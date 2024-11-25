package sqs

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/queue/consumer/backends/backend"
)

type Factory struct {
	logger *zap.Logger
	client *sqs.Client
}

var _ backend.Factory = (*Factory)(nil)

func NewFactory(logger *zap.Logger, client *sqs.Client) *Factory {
	return &Factory{
		logger: logger,
		client: client,
	}
}

// Type implements backend.Factory interface
func (f Factory) Type() backend.Type { return Type }

// Reader implements backend.Factory interface
func (f *Factory) Reader(config *backend.Config) (backend.Backend, error) {
	var (
		backendConfigRaw interface{}
		backendConfig    *Config
		exist            bool
		err              error
	)

	// find config
	if backendConfigRaw, exist = config.Backend[strings.ToLower(string(f.Type()))]; !exist {
		return nil, fmt.Errorf("config for backend consumer not found, %s", f.Type())
	}

	// load raw config to struct
	jsonStr, err := json.Marshal(backendConfigRaw)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(jsonStr, &backendConfig); err != nil {
		return nil, err
	}

	// TODO: need validation ...

	// build new backend instance
	return NewBackend(f.logger, f.client, backendConfig)
}
