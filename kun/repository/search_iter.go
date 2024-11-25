package repository

import (
	"context"
	"database/sql"
	"iter"

	"gorm.io/gorm"

	"github.com/x-challenges/raven/kun/model"
)

type SearchIterator[T model.Model] iter.Seq2[T, error]

// SearchOp interface for repository
type SearchIterOp[T model.Model] interface {
	Operator
	SearchIter(ctx context.Context, filters ...FilterOperator) (SearchIterator[T], error)
}

// Search implement SearchOp interface
type searchIter[T model.Model] struct {
	client func(context.Context) *gorm.DB
}

var _ SearchIterOp[model.Model] = (*searchIter[model.Model])(nil)

func NewSearchIterOp[T model.Model](client func(context.Context) *gorm.DB) SearchIterOp[T] {
	return &searchIter[T]{
		client: client,
	}
}

func (op *searchIter[T]) SearchIter(ctx context.Context, filters ...FilterOperator) (SearchIterator[T], error) {
	operator := &searchIterOperator[T]{
		operatorContext: op,
	}
	return operator.Execute(ctx, filters...)
}

type searchIterOperator[T model.Model] OperatorContext[*searchIter[T]]

func (op *searchIterOperator[T]) Execute(ctx context.Context, filters ...FilterOperator) (SearchIterator[T], error) {
	return func(yield func(T, error) bool) {
		var (
			res  T
			rows *sql.Rows
			err  error
		)

		query := op.operatorContext.client(ctx)

		// apply filters
		query = FilterMerge(query, filters...)

		// take sql rows iterator
		if rows, err = query.Rows(); err != nil {
			yield(res, err)
		}
		defer rows.Close()

		for rows.Next() {
			err = op.operatorContext.client(ctx).ScanRows(rows, &res)

			if !yield(res, err) {
				return
			}
		}
	}, nil
}
