package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/x-challenges/raven/kun"
	"github.com/x-challenges/raven/kun/model"
)

// BatchOp interface for repository
type BatchOp[T model.Model] interface {
	Operator
	Batch(ctx context.Context, ids ...model.ID) ([]T, error)
}

// Batcher implement batch interface
type batch[T model.Model] struct {
	client func(context.Context) *gorm.DB
}

// NewCreate returns create operator
func NewBatchOp[T model.Model](client func(context.Context) *gorm.DB) BatchOp[T] {
	return &batch[T]{
		client: client,
	}
}

// Batch implement Batcher interface
func (op *batch[T]) Batch(ctx context.Context, ids ...model.ID) (result []T, err error) {
	operator := &batchOperator[T]{
		operatorContext: op,
	}

	return operator.Execute(ctx, ids)
}

// batchOperator type provide API for execute query
type batchOperator[T model.Model] OperatorContext[*batch[T]]

// Execute query end encode data to instance
func (op *batchOperator[T]) Execute(ctx context.Context, ids []string) ([]T, error) {
	output := make([]T, 0, len(ids))

	result := op.operatorContext.client(ctx).
		Find(&output, ids)

	return output, kun.HandleError(result.Error)
}
