package dataloaders

import (
	"context"
	"fmt"

	loader "github.com/graph-gophers/dataloader"

	"github.com/x-challenges/raven/common/errors"
)

// Dataloader interface
type Dataloader[T Model] interface {
	// Load instance
	Load(ctx context.Context, key Key) (T, error)

	// LoadMany instances
	LoadMany(ctx context.Context, keys ...Key) ([]T, error)
}

type dataloader[T any] struct {
	loader *loader.Loader
}

type Factory[T Model] func(Fetcher[T]) Dataloader[T]

func New[T Model](opts ...Option) Factory[T] {
	var (
		options = NewOptions()
	)

	// options
	for _, opt := range opts {
		opt(options)
	}

	return func(fetcher Fetcher[T]) Dataloader[T] {
		fn := func(ctx context.Context, keys loader.Keys) []*loader.Result {
			var (
				output = make([]*loader.Result, 0, len(keys))
			)

			// fetch data
			listOf, err := fetcher.Batch(ctx, keys.Keys())

			// transform to map
			objects := sliceToMap(listOf...)

			for _, key := range keys {
				var result *loader.Result

				if err != nil {
					result = &loader.Result{
						Data:  nil,
						Error: err,
					}
					output = append(output, result)

					continue
				}

				obj, ok := objects[key.String()]

				if ok {
					result = &loader.Result{
						Data:  obj,
						Error: nil,
					}
				} else {
					result = &loader.Result{
						Data:  nil,
						Error: errors.ErrNotFound,
					}
				}

				output = append(output, result)
			}
			return output
		}

		return &dataloader[T]{
			loader: loader.NewBatchedLoader(fn, []loader.Option{
				loader.WithBatchCapacity(options.Limit),
				loader.WithClearCacheOnBatch(),
			}...),
		}
	}
}

// Load implement Dataloader interface
func (d *dataloader[T]) Load(ctx context.Context, key Key) (output T, err error) {
	var (
		result any
	)

	if result, err = d.loader.Load(ctx, loader.StringKey(key))(); err != nil {
		return output, err
	}

	return result.(T), nil // TODO: check type assertion ...
}

// LoadMany implement Dataloader interface
func (d *dataloader[T]) LoadMany(ctx context.Context, keys ...Key) ([]T, error) {
	output := make([]T, 0, len(keys))

	result, err := d.loader.LoadMany(ctx, loader.NewKeysFromStrings(keys))()
	if err != nil {
		return nil, fmt.Errorf("dataloader load many failed, %w", errors.Join(err...))
	}

	for _, r := range result {
		output = append(output, r.(T))
	}

	return output, nil
}
