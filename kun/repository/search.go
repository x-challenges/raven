package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/x-challenges/raven/kun"
	"github.com/x-challenges/raven/kun/model"
)

const (
	searchOpDefaultLimit = 30
)

// SearchOp interface for repository
type SearchOp[T model.Model] interface {
	Operator
	Search(ctx context.Context, page Pager, filters ...FilterOperator) ([]T, *PageInfo, error)
}

// Search implement SearchOp interface
type search[T model.Model] struct {
	client func(context.Context) *gorm.DB
}

func NewSearchOp[T model.Model](client func(context.Context) *gorm.DB) SearchOp[T] {
	return &search[T]{
		client: client,
	}
}

// Search implement SearchOp interface
func (op *search[T]) Search(
	ctx context.Context,
	page Pager,
	filters ...FilterOperator,
) ([]T, *PageInfo, error) {
	operator := &searchOperator[T]{
		operatorContext: op,
	}
	return operator.Execute(ctx, page, filters...)
}

type searchOperator[T model.Model] OperatorContext[*search[T]]

// Execute query
func (op *searchOperator[T]) Execute(
	ctx context.Context,
	page Pager,
	filters ...FilterOperator,
) ([]T, *PageInfo, error) {
	var (
		output = make([]T, 0, searchOpDefaultLimit)
	)

	query := op.operatorContext.client(ctx).
		Model(new(T))

	if page != nil {
		filters = append(filters, Page(page))
	}

	// filters
	query = FilterMerge(query, filters...)

	// scan objects
	result := query.Find(&output)

	if err := result.Error; err != nil {
		return nil, nil, kun.HandleError(result.Error)
	}

	start := StartCursor(output)
	end := EndCursor(output)

	var pageInfo = &PageInfo{
		StartCursor: start,
		LastCursor:  end,
		HasNext:     !(end == ""),   // TODO
		HasPrevious: !(start == ""), // TODO
	}

	return output, pageInfo, nil
}
