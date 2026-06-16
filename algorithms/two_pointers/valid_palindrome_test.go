package two_pointers

import "testing"

func TestIsPalindrome(t *testing.T) {
	result := IsPalindrome("racecar")

	if result != true {
		t.Errorf("expected true, got %v", result)
	}
}

func TestIsPalindrome1(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"racecar", true},
		{"hello", false},
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"", true},
		{"a", true},
		{"0P", false},
	}

	for _, tt := range tests {
		result := IsPalindrome(tt.input)

		if result != tt.expected {
			t.Errorf("input=%q expected=%v got=%v",
				tt.input, tt.expected, result)
		}
	}
}
