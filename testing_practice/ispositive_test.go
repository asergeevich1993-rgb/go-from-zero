package main

import "testing"

func isPositive(n int) bool {
	return n >= 0

}

func TestIsPositive(t *testing.T) {
	nums := []struct {
		num  int
		want bool
	}{
		{1, true},
		{-2, false},
		{3, true},
		{0, true}}

	for _, n := range nums {
		got := isPositive(n.num)
		if got != n.want {
			t.Errorf("%d = %v, want %t", n.num, got, n.want)
		}
	}
}
