package consumer

import (
	"context"

	"github.com/x-challenges/raven/modules/queue/consumer/backends/backend"
)

type Message = interface{}

type Resolver interface {
	// Name
	Name() string
}

type Processor interface {
	// Process
	Process(context.Context, ...Message) error
}

type Consumer interface {
	Resolver
	Processor
}

type Wrapper struct {
	Consumer

	backend backend.Backend
}

var _ Consumer = (*Wrapper)(nil)
