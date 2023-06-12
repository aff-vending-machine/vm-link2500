package db

import (
	"reflect"
	"testing"
)

func TestQuery_Where_NoArgs(t *testing.T) {
	q := NewQuery()

	result := q.WhereIf("field = ?")

	if len(result.WhereVal) != 0 {
		t.Errorf("Expected Where slice length to be 0, got %d", len(result.WhereVal))
	}
}

func TestQuery_Where_NilArg(t *testing.T) {
	q := NewQuery()
	var arg *int

	result := q.WhereIf("field = ?", arg)

	if len(result.WhereVal) != 0 {
		t.Errorf("Expected Where slice length to be 0, got %d", len(result.WhereVal))
	}
}

func TestQuery_Where_WithArgs(t *testing.T) {
	q := NewQuery()
	query := "field = ?"
	args := []interface{}{"value"}

	result := q.Where(query, args...)

	if len(result.WhereVal) != 1 {
		t.Errorf("Expected Where slice length to be 1, got %d", len(result.WhereVal))
	}

	where := result.WhereVal[0]
	if where.Query != query {
		t.Errorf("Expected Where query to be %q, got %q", query, where.Query)
	}

	if !reflect.DeepEqual(where.Args, args) {
		t.Errorf("Expected Where args to be %v, got %v", args, where.Args)
	}
}

func TestQuery_WhereIf_NoArgs(t *testing.T) {
	q := NewQuery()

	result := q.WhereIf("field = ?")

	if len(result.WhereVal) != 0 {
		t.Errorf("Expected Where slice length to be 0, got %d", len(result.WhereVal))
	}
}

func TestQuery_WhereIf_NilArg(t *testing.T) {
	q := NewQuery()
	var arg *int

	result := q.WhereIf("field = ?", arg)

	if len(result.WhereVal) != 0 {
		t.Errorf("Expected Where slice length to be 0, got %d", len(result.WhereVal))
	}
}

func TestQuery_WhereIf_WithArgs(t *testing.T) {
	q := NewQuery()
	query := "field = ?"
	args := []interface{}{"value"}

	result := q.WhereIf(query, args...)

	if len(result.WhereVal) != 1 {
		t.Errorf("Expected Where slice length to be 1, got %d", len(result.WhereVal))
	}

	where := result.WhereVal[0]
	if where.Query != query {
		t.Errorf("Expected Where query to be %q, got %q", query, where.Query)
	}

	if !reflect.DeepEqual(where.Args, args) {
		t.Errorf("Expected Where args to be %v, got %v", args, where.Args)
	}
}

func TestQuery_OrderIf_NilOrder(t *testing.T) {
	q := NewQuery()

	result := q.OrderIf(nil)

	if result.OrderVal != nil {
		t.Error("Expected Order to be nil")
	}
}

func TestQuery_OrderIf_WithOrder(t *testing.T) {
	q := NewQuery()
	order := "field:desc"

	result := q.OrderIf(&order)

	expectedField := "field"
	expectedDescending := true
	if result.OrderVal == nil {
		t.Error("Expected Order to be non-nil")
	} else {
		if result.OrderVal.Field != expectedField {
			t.Errorf("Expected Order Field to be %q, got %q", expectedField, result.OrderVal.Field)
		}
		if result.OrderVal.Decending != expectedDescending {
			t.Errorf("Expected Order Decending to be %t, got %t", expectedDescending, result.OrderVal.Decending)
		}
	}
}

func TestQuery_LimitIfNotNil(t *testing.T) {
	q := NewQuery()
	limit := 10

	result := q.LimitIfNotNil(&limit)

	if result.LimitVal == nil {
		t.Error("Expected Limit to be non-nil")
	} else if *result.LimitVal != limit {
		t.Errorf("Expected Limit to be %d, got %d", limit, *result.LimitVal)
	}
}

func TestQuery_OffsetIf(t *testing.T) {
	q := NewQuery()
	offset := 5

	result := q.OffsetIf(&offset)

	if result.OffsetVal == nil {
		t.Error("Expected Offset to be non-nil")
	} else if *result.OffsetVal != offset {
		t.Errorf("Expected Offset to be %d, got %d", offset, *result.OffsetVal)
	}
}

func TestQuery_PreloadsIf_NoPreloads(t *testing.T) {
	q := NewQuery()

	result := q.PreloadsIf(nil)

	if len(result.PerloadVal) != 0 {
		t.Errorf("Expected Perloads slice length to be 0, got %d", len(result.PerloadVal))
	}
}

func TestQuery_PreloadsIf_WithPreloads(t *testing.T) {
	q := NewQuery()
	preloads := "preload1:preload2"

	result := q.PreloadsIf(&preloads)

	expectedPreloads := []preload{
		{Query: "preload1"},
		{Query: "preload2"},
	}

	if len(result.PerloadVal) != len(expectedPreloads) {
		t.Errorf("Expected Perloads slice length to be %d, got %d", len(expectedPreloads), len(result.PerloadVal))
	}

	for i, preload := range result.PerloadVal {
		expectedPreload := expectedPreloads[i]
		if preload.Query != expectedPreload.Query {
			t.Errorf("Expected Preload query to be %q, got %q", expectedPreload.Query, preload.Query)
		}
	}
}

func TestQuery_Clear(t *testing.T) {
	limit := 10
	offset := 5
	q := NewQuery()
	q.WhereIf("field = ?", "value")
	q.LimitIfNotNil(&limit)
	q.OffsetIf(&offset)
	q.Preload("preload")

	result := q.Clear()

	if len(result.WhereVal) != 0 {
		t.Errorf("Expected Where slice length to be 0, got %d", len(result.WhereVal))
	}

	if result.LimitVal != nil {
		t.Error("Expected Limit to be nil")
	}

	if result.OffsetVal != nil {
		t.Error("Expected Offset to be nil")
	}

	if len(result.PerloadVal) != 0 {
		t.Errorf("Expected Perloads slice length to be 0, got %d", len(result.PerloadVal))
	}
}

func TestQuery_Limit(t *testing.T) {
	q := NewQuery()
	limit := 10

	result := q.Limit(limit)

	if result.LimitVal == nil {
		t.Error("Expected Limit to be non-nil")
	} else if *result.LimitVal != limit {
		t.Errorf("Expected Limit to be %d, got %d", limit, *result.LimitVal)
	}
}

func TestQuery_Offset(t *testing.T) {
	q := NewQuery()
	offset := 5

	result := q.Offset(offset)

	if result.OffsetVal == nil {
		t.Error("Expected Offset to be non-nil")
	} else if *result.OffsetVal != offset {
		t.Errorf("Expected Offset to be %d, got %d", offset, *result.OffsetVal)
	}
}

func TestQuery_Order(t *testing.T) {
	q := NewQuery()
	order := "field:desc"

	result := q.Order(order)

	expectedField := "field"
	expectedDescending := true
	if result.OrderVal == nil {
		t.Error("Expected Order to be non-nil")
	} else {
		if result.OrderVal.Field != expectedField {
			t.Errorf("Expected Order Field to be %q, got %q", expectedField, result.OrderVal.Field)
		}
		if result.OrderVal.Decending != expectedDescending {
			t.Errorf("Expected Order Decending to be %t, got %t", expectedDescending, result.OrderVal.Decending)
		}
	}
}
