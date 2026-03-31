package main

import "testing"

func TestPrintResult(t *testing.T) {
	tests := []struct {
		result float64
		err error
		want float64
	}{
		{1, nil, 1},
		{2, nil, 2},
	}

	for _, tt := range tests {
		got, err := printResult(tt.result, tt.err)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if got != tt.want {
			t.Errorf("expected %f, got %f", tt.want, got)
		}
	}
}