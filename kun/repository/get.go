package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/x-challenges/raven/kun"
	"github.com/x-challenges/raven/kun/model"
)

// GetOp public interface
type GetOp[T model.Model] interface {
	// Get
	Get(ctx context.Context, id model.ID) (T, error)
}

// GetOp private interface
type getOp[T model.Model, PT ptr[T]] interface {
	Operator
	GetOp[PT]
}

// get implement Getter interface
type get[T model.Model, PT ptr[T]] struct {
	client func(context.Context) *gorm.DB
}

// NewGetOp returns new get operator
func NewGetOp[PT ptr[T], T model.Model](client func(context.Context) *gorm.DB) GetOp[PT] {
	return &get[T, PT]{
		client: client,
	}
}

// Get implement Getter interface
func (op *get[T, PT]) Get(ctx context.Context, id model.ID) (PT, error) {
	operator := &getOperator[T, PT]{
		operatorContext: op,
	}
	return operator.Execute(ctx, id)
}

// getOperator type provide API for execute query
type getOperator[T model.Model, PT ptr[T]] OperatorContext[*get[T, PT]]

func (op *getOperator[T, PT]) Execute(ctx context.Context, id model.ID) (PT, error) {
	instance := new(T)

	result := op.operatorContext.client(ctx).
		First(instance, "id = ?", id)

	return instance, kun.HandleError(result.Error)
}
