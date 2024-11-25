package dataloaders

import "context"

type Key = string

type Model interface {
	GetID() string
}

type Fetcher[T Model] interface {
	Batch(context.Context, []string) ([]T, error)
}
