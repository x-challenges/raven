package repository

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type Filtered interface{}

type FilteredCommon interface {
	Filtered
	~int | ~int32 | ~string | ~bool
}

type FilteredString interface {
	Filtered
	~string
}

type FilterOperator func(query *gorm.DB) *gorm.DB

type Filter[T FilteredCommon] func(field string, value T) FilterOperator

func Empty[T FilteredCommon]() FilterOperator {
	return func(query *gorm.DB) *gorm.DB {
		return query
	}
}

// IsNull filter
func IsNull(field string) FilterOperator {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where(field + " IS NULL")
	}
}

// Equal filter
func Equal[T FilteredCommon](field string, value T) FilterOperator {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where(field+" = ?", value)
	}
}

// NotEqual filter
func NotEqual[T FilteredCommon](field string, value T) FilterOperator {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where(field+" != ?", value)
	}
}

// In filter
func In[T FilteredCommon](field string, values []T) FilterOperator {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where(field+" IN (?)", values)
	}
}

// Has Json filter
func HasJSON[T FilteredCommon](field string, value T) FilterOperator {
	return func(query *gorm.DB) *gorm.DB {
		// TODO: gotm has some problems with ? and @ escaping ...
		return query.Where(`JSON_EXISTS(` + field + `, '$[*] ? (@ == "` + fmt.Sprintf("%v", value) + `")')`)
	}
}

// InOr filter
func InOr[T FilteredCommon](field string, values []T) FilterOperator {
	return func(query *gorm.DB) *gorm.DB {
		return query.Or(field+" IN (?)", values)
	}
}

// NotIn filter
func NotIn[T FilteredCommon](field string, values []T) FilterOperator {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where(field+" NOT IN (?)", values)
	}
}

// And group filter
func And(ops ...FilterOperator) FilterOperator {
	return func(query *gorm.DB) *gorm.DB {
		var (
			merged *gorm.DB = query
		)

		for _, op := range ops {
			merged = FilterMerge(merged, op)
		}

		return query.Where(merged)
	}
}

// Or group filter
func Or(ops ...FilterOperator) FilterOperator {
	return func(query *gorm.DB) *gorm.DB {
		var (
			merged *gorm.DB = query
		)

		for _, op := range ops {
			merged = FilterMerge(merged, op)
		}

		return query.Or(merged)

		// return query.Or(" OR ", func(q *gorm.DB) *gorm.DB {
		// 	for _, op := range ops {
		// 		q = FilterMerge(q, op)
		// 	}
		// 	return q
		// })
	}
}

// Like filter
func Like[T FilteredString](field string, value T) FilterOperator {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where(field+" LIKE ?", value+"%")
	}
}

// LikeInsens filter
func LikeInsens[T FilteredString](field string, value T) FilterOperator {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where(field+" ILIKE ?", value+"%")
	}
}

// Page filter
func Page(page Pager) FilterOperator {
	return func(query *gorm.DB) *gorm.DB {
		if page != nil {
			// limit offset
			if after := page.GetAfter(); after > 0 {
				query = query.Limit(int(after))
			}

			// cursor
			if first := page.GetFirst(); first != "" {
				query = query.Where("id > ?", first)
			}
		}
		return query
	}
}

// Order
func Order(fields ...string) FilterOperator {
	return func(query *gorm.DB) *gorm.DB {
		return query.Order(strings.Join(fields, ","))
	}
}

// Limit filter
func Limit(limit int) FilterOperator {
	return func(query *gorm.DB) *gorm.DB {
		return query.Limit(limit)
	}
}

func FilterMerge(query *gorm.DB, filters ...FilterOperator) *gorm.DB {
	for _, filter := range filters {
		if filter != nil {
			query = filter(query)
		}
	}
	return query
}
