package stores

import (
	"fmt"

	"github.com/eko/gocache/lib/v4/store"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type kind string

type Store interface {
	store.StoreInterface
}

type Options struct {
	Memory *memoryOptions `mapstructure:"memory"`
	Redis  *redisOptions  `mapstructure:"redis"`
}

type builder func(options *Options) (Store, error)

type factory func() (kind, builder)

type Factory struct {
	logger   *zap.Logger
	builders map[kind]builder
}

type NewFactoryParamsFx struct {
	fx.In

	Logger    *zap.Logger
	Factories []factory `group:"cache:stores:factory"`
}

func NewFactory(p NewFactoryParamsFx) *Factory {
	var (
		builders = make(map[kind]builder)
	)

	for _, f := range p.Factories {
		kind, builder := f()
		builders[kind] = builder
	}

	return &Factory{
		logger:   p.Logger,
		builders: builders,
	}
}

// Store
func (f *Factory) Store(options *Options) (Store, error) {
	var (
		kind    kind
		builder builder
		exist   bool
	)

	if options == nil {
		options = &Options{
			Memory: &memoryOptions{},
		}
	}

	switch {
	case options.Redis != nil:
		kind = Redis
	case options.Memory == nil:
		fallthrough
	default:
		kind = Memory
		options = &Options{
			Memory: &memoryOptions{},
		}
	}

	if builder, exist = f.builders[kind]; exist {
		return builder(options)
	}

	return nil, fmt.Errorf("unexpected cache type or options, %s", kind)
}
