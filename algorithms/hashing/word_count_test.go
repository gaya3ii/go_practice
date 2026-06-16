package hashing

import "testing"

func TestCharCount(t *testing.T) {
	got := WordCount("golang")
	want := 2
	if got["g"] != 2 {
		// 3. if not — tell t to report failure
		t.Errorf("expected g:%d, got%d", want, got["g"])
	}

}

func TestCharCount1(t *testing.T) {

	tests := []struct {
		input string
		char  string
		want  int
	}{
		{"golang", "g", 2},
		{"golang", "o", 1},
		{"golang", "a", 1},
		{"hello", "l", 2},
	}

	for _, tt := range tests {
		got := WordCount(tt.input)
		if got[tt.char] != tt.want {
			t.Errorf("charCount(%q)[%q] = %d, want %d", tt.input, tt.char, got[tt.char], tt.want)
		}
	}
}
