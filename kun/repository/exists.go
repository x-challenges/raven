package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/x-challenges/raven/kun"
	"github.com/x-challenges/raven/kun/model"
)

// ExistsOp interface for repository
type ExistsOp[T model.Model] interface {
	Operator
	Exists(ctx context.Context, filters ...FilterOperator) (bool, error)
}

// Exists implement ExistsOp interface
type exists[T model.Model] struct {
	client func(context.Context) *gorm.DB
}

func NewExistsOp[T model.Model](client func(context.Context) *gorm.DB) ExistsOp[T] {
	return &exists[T]{
		client: client,
	}
}

// Exists implement ExistsOp interface
func (o *exists[T]) Exists(ctx context.Context, filters ...FilterOperator) (bool, error) {
	operator := &existsOperator[T]{
		operatorContext: o,
	}

	return operator.Execute(ctx, filters...)
}

type existsOperator[T model.Model] OperatorContext[*exists[T]]

func (op *existsOperator[T]) Execute(ctx context.Context, filters ...FilterOperator) (bool, error) {
	query := op.operatorContext.client(ctx).
		Model(new(T))

	var count int64

	result := FilterMerge(query, filters...).Count(&count)

	return count > 0, kun.HandleError(result.Error)
}
