package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/x-challenges/raven/kun"
	"github.com/x-challenges/raven/kun/model"
)

// UpdateOp interface for repository
type UpdateOp[T model.Model] interface {
	Operator
	Update(ctx context.Context, update T, fields ...string) error
}

// Update implement Updater interface
type update[T model.Model] struct {
	client func(context.Context) *gorm.DB
}

func NewUpdateOp[T model.Model](client func(context.Context) *gorm.DB) UpdateOp[T] {
	return &update[T]{
		client: client,
	}
}

// Update implement Updater interface
func (op *update[T]) Update(ctx context.Context, update T, fields ...string) error {
	operator := &updateOperator[T]{
		operatorContext: op,
	}
	return operator.Execute(ctx, update, fields...)
}

type updateOperator[T model.Model] OperatorContext[*update[T]]

// Execute query
func (op *updateOperator[T]) Execute(ctx context.Context, update T, fields ...string) error {
	var query = op.operatorContext.client(ctx).
		Model(update)

	// partial updates
	if len(fields) > 0 {
		query = query.Select(fields)
	} else {
		query = query.Select("*")
	}

	var result = query.Updates(update)

	return kun.HandleError(result.Error)
}
