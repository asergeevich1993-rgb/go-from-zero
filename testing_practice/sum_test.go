package main

import "testing"

func sum1(a, b int) int {
	return a + b

}
func TestSum1(t *testing.T) {
	sum := []struct {
		a, b int
		want int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{4, 3, 7},
		{5, 5, 10}}

	for _, s := range sum {
		got := sum1(s.a, s.b)
		if got != s.want {
			t.Errorf("Sum %d + %d  равна %v, want %d", s.a, s.b, got, s.want)
		}
	}
}
