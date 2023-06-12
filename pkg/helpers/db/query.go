package db

import (
	"reflect"
	"strings"
)

type Query struct {
	LimitVal   *int
	OffsetVal  *int
	WhereVal   []where
	OrderVal   *order
	PerloadVal []preload // gorm only
}

type where struct {
	Query string
	Args  []interface{}
}

type order struct {
	Field     string
	Decending bool
}

type preload struct {
	Query string
	Args  []interface{}
}

func NewQuery() *Query {
	return &Query{}
}

// Where appends a where clause
func (q *Query) Where(query string, args ...interface{}) *Query {
	q.WhereVal = append(q.WhereVal, where{Query: query, Args: args})
	return q
}

// Preload appends a preload
func (q *Query) Preload(query string, args ...interface{}) *Query {
	q.PerloadVal = append(q.PerloadVal, preload{Query: query, Args: args})
	return q
}

// Limit sets the limit
func (q *Query) Limit(limit int) *Query {
	q.LimitVal = &limit
	return q
}

// Offset sets the offset
func (q *Query) Offset(offset int) *Query {
	q.OffsetVal = &offset
	return q
}

// Order sets the order
func (q *Query) Order(sort string) *Query {
	decending := strings.Contains(sort, "desc")
	q.OrderVal = &order{Field: strings.TrimSuffix(sort, ":desc"), Decending: decending}
	return q
}

func (q *Query) Clear() *Query {
	q.WhereVal = []where{}
	q.LimitVal = nil
	q.OffsetVal = nil
	q.PerloadVal = []preload{}
	return q
}

// WhereIf adds a where clause if the pointer is not nil
func (q *Query) WhereIf(query string, args ...interface{}) *Query {
	// Return early if there are no arguments
	if len(args) == 0 {
		return q
	}

	// Use a type switch to handle different types
	value := reflect.ValueOf(args[0])
	if value.Kind() == reflect.Ptr && value.IsNil() {
		return q
	}

	// Add the where clause
	q.Where(query, args...)
	return q
}

// OrderIf sets the order if the pointer is not nil
func (q *Query) OrderIf(sort *string) *Query {
	if sort == nil {
		return q
	}

	q.Order(*sort)
	return q
}

// LimitIfNotNil sets the limit if the pointer is not nil
func (q *Query) LimitIfNotNil(limit *int) *Query {
	q.LimitVal = limit
	return q
}

// OffsetIf sets the offset if the pointer is not nil
func (q *Query) OffsetIf(offset *int) *Query {
	q.OffsetVal = offset
	return q
}

// PreloadsIf adds preloads if the pointer is not nil
func (q *Query) PreloadsIf(preloads *string) *Query {
	// Return early if there are no arguments
	if preloads == nil {
		return q
	}

	list := strings.Split(*preloads, ":")
	for _, preload := range list {
		q.Preload(preload)
	}

	// Add the where clause
	return q
}
