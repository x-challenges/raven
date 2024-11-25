package consumer

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/queue/consumer/backends/backend"
)

type Dispatcher interface {
	// Run
	Run(context.Context) error

	// Shutdown
	Shutdown(context.Context) error

	// Register
	Register(consumer Consumer) error
}

type dispatcher struct {
	logger    *zap.Logger
	config    *Config
	factories map[backend.Type]backend.Factory
	consumers map[string]Wrapper
}

type DispatcherParams struct {
	fx.In

	Logger    *zap.Logger
	Config    *Config
	Factories []backend.Factory `group:"queue:consumer:backend:factory"`
}

func NewDispatcher(p DispatcherParams) Dispatcher {
	var factories = map[backend.Type]backend.Factory{}

	for _, factory := range p.Factories {
		factories[strings.ToLower(factory.Type())] = factory
	}

	return &dispatcher{
		logger:    p.Logger,
		config:    p.Config,
		consumers: make(map[string]Wrapper),
		factories: factories,
	}
}

// Register implements Dispatcher interface
func (d *dispatcher) Register(cons Consumer) error {
	var (
		consumerName   = cons.Name()
		backendType    backend.Type
		backendConfig  backend.Config
		backendFactory backend.Factory
		backendImpl    backend.Backend
		ok             bool
		err            error
	)

	// resolve backend type from config
	var exist bool

	for name, v := range d.config.Queue.Consumers {
		if name == consumerName {
			exist = true

			if len(v.Backend) != 1 {
				return fmt.Errorf("consumer must have only one backend configurration, %s", consumerName)
			}

			backendConfig = v.Config

			for bt := range v.Backend {
				backendType = backend.Type(bt)
			}

			break
		}
	}

	if !exist {
		return fmt.Errorf("dont have configuration for consumer with name, %s", consumerName)
	}

	// resolve backend factory
	if backendFactory, ok = d.factories[backendType]; !ok {
		return fmt.Errorf("unexpected consumer backend type, %s", backendType)
	}

	// build backend
	if backendImpl, err = backendFactory.Reader(&backendConfig); err != nil {
		return err
	}

	// register consumer
	d.consumers[consumerName] = Wrapper{
		Consumer: cons,
		backend:  backendImpl,
	}

	return nil
}

// Run implements Dispatcher interface
func (d *dispatcher) Run(ctx context.Context) error {
	for _, c := range d.consumers {
		if err := c.backend.Run(ctx, c.Process); err != nil {
			return err
		}
	}
	return nil
}

// Shutdown implements Disptcher interface
func (d *dispatcher) Shutdown(ctx context.Context) error {
	var errs []error

	for _, c := range d.consumers {
		if err := c.backend.Close(ctx); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}
