package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/x-challenges/raven/kun"
	"github.com/x-challenges/raven/kun/model"
)

// FindOp interface for repository
type FindOp[T model.Model] interface {
	Operator
	Find(ctx context.Context, filters ...FilterOperator) (T, error)
}

// find implement FindOp interface
type find[T model.Model, PT ptr[T]] struct {
	client func(context.Context) *gorm.DB
}

// NewFindOp returns new find operator
func NewFindOp[PT ptr[T], T model.Model](client func(context.Context) *gorm.DB) FindOp[PT] {
	return &find[T, PT]{
		client: client,
	}
}

// Find implement FindOp interface
func (op *find[T, PT]) Find(ctx context.Context, filters ...FilterOperator) (PT, error) {
	operator := &findOperator[T, PT]{
		operatorContext: op,
	}

	return operator.Execute(ctx, filters...)
}

// findOperator type provide API for execute query
type findOperator[T model.Model, PT ptr[T]] OperatorContext[*find[T, PT]]

func (op *findOperator[T, PT]) Execute(ctx context.Context, filters ...FilterOperator) (PT, error) {
	instance := new(T)

	query := op.operatorContext.client(ctx)

	result := FilterMerge(query, filters...).
		Find(instance)

	if err := kun.HandleError(result.Error); err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, kun.HandleError(gorm.ErrRecordNotFound)
	}

	return instance, nil
}
