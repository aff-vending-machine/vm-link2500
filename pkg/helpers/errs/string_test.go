package errs

import (
	"errors"
	"testing"
)

type StringTestcase struct {
	name     string
	err      error
	expected string
}

func TestErrorToString(t *testing.T) {
	err := errors.New("test error")

	tests := []StringTestcase{
		{
			name:     "non-nil error",
			err:      err,
			expected: "test error",
		},
		{
			name:     "nil error",
			err:      nil,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToString(tt.err)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}
