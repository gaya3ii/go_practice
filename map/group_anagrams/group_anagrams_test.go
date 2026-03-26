package main

import "testing"

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, 3},
		{[]string{"a"}, 1},
		{[]string{""}, 1},
	}

	for _, tt := range tests {
		got := groupAnagrams(tt.input)
		if len(got) != tt.want {
			t.Errorf("groupAnagrams(%v) = %d groups, want %d", tt.input, len(got), tt.want)
		}
	}
}
