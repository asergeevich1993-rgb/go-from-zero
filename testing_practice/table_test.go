package main

import "testing"

func isEven1(n int) bool {
	return n%2 == 0
}

func TestIsEven1Table(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{2, true},
		{3, false},
		{0, true},
		{-1, false},
		{100, true}}
	for _, tt := range tests {
		got := isEven(tt.input)
		if got != tt.want {
			t.Errorf("isEven(%d)=%v;want %v", tt.input, got, tt.want)
		}
	}
}
