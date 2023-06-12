package gen

import "testing"

type Testcase struct {
	name     string
	length   int
	expected int
}

func TestGenerateRandom(t *testing.T) {
	tests := []Testcase{
		{
			name:     "generate random string of length 10",
			length:   10,
			expected: 10,
		},
		{
			name:     "generate random string of length 20",
			length:   20,
			expected: 20,
		},
		{
			name:     "generate random string of length 0",
			length:   0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Random(tt.length)
			if len(result) != tt.expected {
				t.Errorf("expected length of %d, got %d", tt.expected, len(result))
			}
			for _, char := range result {
				if (char < '0' || char > '9') && (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') {
					t.Errorf("expected alphanumeric characters, got %q", char)
				}
			}
		})
	}
}

func TestGenerateit(t *testing.T) {
	str := UUIDv4()
	expectedLength := 36
	if len(str) != expectedLength {
		t.Errorf("expected length of %d, got %d", expectedLength, len(str))
	}

	expectedFormat := "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx"
	if len(str) != len(expectedFormat) {
		t.Errorf("expected format length of %d, got %d", len(expectedFormat), len(str))
	}

	// Check that the format contains expected characters
	for i, char := range str {
		expectedChar := expectedFormat[i]
		if expectedChar != 'x' && expectedChar != 'y' {
			continue
		}
		if expectedChar == 'x' && char == '-' {
			t.Errorf("expected 'x' at position %d, got '-'", i)
		}
		if expectedChar == 'y' && char != '8' && char != '9' && char != 'a' && char != 'b' {
			t.Errorf("expected 'y' at position %d to be '8', '9', 'a', or 'b', got %q", i, char)
		}
	}
}

func TestGenerateOTP(t *testing.T) {
	str := OTP()
	expectedLength := 6
	if len(str) != expectedLength {
		t.Errorf("expected length of %d, got %d", expectedLength, len(str))
	}

	for _, char := range str {
		if char < '0' || char > '9' {
			t.Errorf("expected only digits, got %q", char)
		}
	}
}
