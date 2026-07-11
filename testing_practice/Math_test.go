package main

import "testing"

func sum(a, b int) int {
	return a + b
}

func TestSum(t *testing.T) {
	result := sum(2, 3)
	if result != 5 {
		t.Errorf("sum(2,3)=%d;want 5", result)
	}
}
