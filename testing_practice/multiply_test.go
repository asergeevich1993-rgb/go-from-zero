package main

import "testing"

func multiply(a, b int) int {
	return a * b
}

func TestMultiply(t *testing.T) {

	nums := []struct {
		a, b int
		want int
	}{
		{3, 4, 12},
		{5, 5, 25},
		{7, 7, 49},
		{0, 7, 0}}

	for _, n := range nums {
		got := multiply(n.a, n.b)
		if got != n.want {
			t.Errorf("%d*%d=%v,want %d", n.a, n.b, got, n.want)

		}
	}
}
