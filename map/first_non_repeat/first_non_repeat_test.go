package main

import "testing"

func TestNonRepeat(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"leetcode", "l"},
		{"loveleetcode", "v"},
		{"aabb", ""},
	}
	for _, tt := range tests {
		got := nonRepeat(tt.input)
		if got != tt.want {
			t.Errorf("expected %s, got %s", tt.want, got)
		}
	}
}
