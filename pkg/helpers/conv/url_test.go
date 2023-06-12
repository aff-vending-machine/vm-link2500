package conv

import (
	"testing"
)

func TestToURLPath(t *testing.T) {
	tests := []struct {
		name     string
		params   []string
		expected string
	}{
		{
			name:     "should handle normal parameters",
			params:   []string{"api", "v1", "users"},
			expected: "api/v1/users",
		},
		{
			name:     "should trim trailing and leading slashes",
			params:   []string{"/api/", "/v1/", "/users/"},
			expected: "api/v1/users",
		},
		{
			name:     "should handle empty parameters",
			params:   []string{"", "", ""},
			expected: "",
		},
		{
			name:     "should handle mixed parameters",
			params:   []string{"api", "/v1", "users/"},
			expected: "api/v1/users",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := JoinURLPathSegments(tt.params...)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}
