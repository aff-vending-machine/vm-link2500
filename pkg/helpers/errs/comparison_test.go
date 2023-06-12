package errs

import (
	"errors"
	"testing"
)

type ComparisonTestcase struct {
	name     string
	err      error
	msg      string
	target   error
	expected bool
}

func TestHasMsg(t *testing.T) {
	tests := []ComparisonTestcase{
		{
			name:     "error contains message",
			err:      errors.New("this is a sample error"),
			msg:      "sample",
			expected: true,
		},
		{
			name:     "error does not contain message",
			err:      errors.New("this is another error"),
			msg:      "sample",
			expected: false,
		},
		{
			name:     "error is nil",
			err:      nil,
			msg:      "sample",
			expected: false,
		},
		{
			name:     "case insensitive check",
			err:      errors.New("this is an Error"),
			msg:      "error",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HasMsg(tt.err, tt.msg)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}

func TestNoMsg(t *testing.T) {
	tests := []ComparisonTestcase{
		{
			name:     "error contains message",
			err:      errors.New("this is a sample error"),
			msg:      "sample",
			expected: false,
		},
		{
			name:     "error does not contain message",
			err:      errors.New("this is another error"),
			msg:      "sample",
			expected: true,
		},
		{
			name:     "error is nil",
			err:      nil,
			msg:      "sample",
			expected: true,
		},
		{
			name:     "case insensitive check",
			err:      errors.New("this is an Error"),
			msg:      "error",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NoMsg(tt.err, tt.msg)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}

func TestIs(t *testing.T) {
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	targetErr := errors.New("target error")

	tests := []ComparisonTestcase{
		{
			name:     "same error",
			err:      err1,
			target:   err1,
			expected: true,
		},
		{
			name:     "different errors",
			err:      err1,
			target:   err2,
			expected: false,
		},
		{
			name:     "target error is nil",
			err:      err1,
			target:   nil,
			expected: false,
		},
		{
			name:     "errors.Is check",
			err:      targetErr,
			target:   targetErr,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Is(tt.err, tt.target)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}

func TestNot(t *testing.T) {
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	targetErr := errors.New("target error")

	tests := []ComparisonTestcase{
		{
			name:     "same error",
			err:      err1,
			target:   err1,
			expected: false,
		},
		{
			name:     "different errors",
			err:      err1,
			target:   err2,
			expected: true,
		},
		{
			name:     "target error is nil",
			err:      err1,
			target:   nil,
			expected: true,
		},
		{
			name:     "errors.Is check",
			err:      targetErr,
			target:   targetErr,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Not(tt.err, tt.target)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}
