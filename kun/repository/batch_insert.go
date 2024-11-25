package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/x-challenges/raven/kun"
	"github.com/x-challenges/raven/kun/model"
)

// BatchInsertOp interface for repository
type BatchInsertOp[T model.Model] interface {
	Operator
	BatchInsert(ctx context.Context, instances ...T) error
}

// BatchInsert implement BatchInsertOp interface
type batchInsert[T model.Model] struct {
	client func(context.Context) *gorm.DB
}

// NewBatchInsert returns BatchInsert operator
func NewBatchInsertOp[T model.Model](client func(context.Context) *gorm.DB) BatchInsertOp[T] {
	return &batchInsert[T]{
		client: client,
	}
}

// BatchInsert implement BatchInsertOp interface
func (o *batchInsert[T]) BatchInsert(ctx context.Context, instances ...T) error {
	operator := &batchInsertOperator[T]{
		operatorContext: o,
	}

	return operator.Execute(ctx, instances...)
}

// BatchInsertOperator type provide API for execute query
type batchInsertOperator[T model.Model] OperatorContext[*batchInsert[T]]

// Execute query
func (op *batchInsertOperator[T]) Execute(ctx context.Context, instances ...T) error {
	result := op.operatorContext.client(ctx).
		CreateInBatches(instances, 100)

	return kun.HandleError(result.Error)
}
