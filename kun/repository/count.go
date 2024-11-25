package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/x-challenges/raven/kun"
)

// CountOp public interface
type CountOp[T any] interface {
	Count(ctx context.Context, filters ...FilterOperator) (int, error)
}

// CountOp private interface
type countOp[T any, PT ptr[T]] interface {
	Operator
	CountOp[T]
}

// Count implement CountOp interface
type count[T any, PT ptr[T]] struct {
	client func(context.Context) *gorm.DB
}

func NewCountOp[PT ptr[T], T any](client func(context.Context) *gorm.DB) CountOp[PT] {
	return &count[T, PT]{
		client: client,
	}
}

// Count implement CountOp interface
func (o *count[T, PT]) Count(ctx context.Context, filters ...FilterOperator) (int, error) {
	operator := &countOperator[T, PT]{
		operatorContext: o,
	}

	return operator.Execute(ctx, filters...)
}

type countOperator[T any, PT ptr[T]] OperatorContext[*count[T, PT]]

func (op *countOperator[T, PT]) Execute(ctx context.Context, filters ...FilterOperator) (int, error) {
	var (
		output int64
	)

	query := op.operatorContext.client(ctx).
		Model(new(T))

	result := FilterMerge(query, filters...).
		Count(&output)

	return int(output), kun.HandleError(result.Error)
}
