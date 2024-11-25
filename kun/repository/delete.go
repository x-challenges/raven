package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/x-challenges/raven/kun"
	"github.com/x-challenges/raven/kun/model"
)

// DeleteOp public interface
type DeleteOp[T any] interface {
	Delete(ctx context.Context, id model.ID) error
}

// DeleteOp private interface
type deleteOp[T any, PT ptr[T]] interface {
	Operator
	DeleteOp[T]
}

// Delete implement DeleteOp interface
type delete[T any, PT ptr[T]] struct {
	client func(context.Context) *gorm.DB
}

func NewDeleteOp[PT ptr[T], T any](client func(context.Context) *gorm.DB) DeleteOp[PT] {
	return &delete[T, PT]{
		client: client,
	}
}

// Delete implement DeleteOp interface
func (o *delete[T, PT]) Delete(ctx context.Context, id model.ID) error {
	operator := &deleteOperator[T, PT]{
		operatorContext: o,
	}

	return operator.Execute(ctx, id)
}

type deleteOperator[T any, PT ptr[T]] OperatorContext[*delete[T, PT]]

func (op *deleteOperator[T, PT]) Execute(ctx context.Context, id model.ID) error {
	result := op.operatorContext.client(ctx).
		Where("id = ?", id).
		Delete(new(T))

	return kun.HandleError(result.Error)
}
