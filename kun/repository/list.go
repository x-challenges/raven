package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/x-challenges/raven/kun"
	"github.com/x-challenges/raven/kun/model"
)

type ListOp[T model.Model] interface {
	Operator

	List(ctx context.Context, filters ...FilterOperator) ([]T, error)
}

type list[T model.Model] struct {
	client func(context.Context) *gorm.DB
}

func NewListOp[T model.Model](client func(context.Context) *gorm.DB) ListOp[T] {
	return &list[T]{
		client: client,
	}
}

func (op *list[T]) List(ctx context.Context, filters ...FilterOperator) ([]T, error) {
	operator := &listOperator[T]{
		operatorContext: op,
	}
	return operator.Execute(ctx, filters...)
}

type listOperator[T model.Model] OperatorContext[*list[T]]

func (op *listOperator[T]) Execute(ctx context.Context, filters ...FilterOperator) ([]T, error) {
	var (
		output = []T{}
	)

	var query = op.operatorContext.client(ctx).
		Model(new(T))

	// filters
	query = FilterMerge(query, filters...)

	// scan objects
	result := query.Find(&output)

	if err := result.Error; err != nil {
		return nil, kun.HandleError(result.Error)
	}

	return output, nil
}
