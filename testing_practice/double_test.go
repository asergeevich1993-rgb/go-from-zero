package main

import (
	"reflect"
	"testing"
)

func double(num []int) []int {
	double := []int{}
	for _, n := range num {
		if n == 0 {

			return double
		}
		n = n * 2
		double = append(double, n)
	}
	return double

}
func TestDouble(t *testing.T) {
	num := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{2, 4, 6, 8, 10}},
		{[]int{10, 20, 30, 40}, []int{20, 40, 60, 80}},
		{[]int{0}, []int{}}}

	for _, n := range num {
		got := double(n.nums)
		if !reflect.DeepEqual(got, n.want) {
			t.Errorf("%v = %v, want %v", n.nums, got, n.want)
		}
	}
}
