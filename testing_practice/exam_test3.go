package main

import (
	"reflect"
	"testing"
)

func filterbyeven(nums []int) []int {
	even := []int{}
	for _, n := range nums {
		if n%2 == 0 {
			even = append(even, n)
		}
	}
	return even

}

func TestFilterbyeven(t *testing.T) {
	nums := []struct {
		num  []int
		want []int
	}{
		{[]int{2, 3, 4, 5, 6}, []int{2, 4, 6}},
		{[]int{10, 8, 7, 2, 1}, []int{10, 8, 2}},
		{[]int{0}, []int{0}}}

	for _, n := range nums {
		got := filterbyeven(n.num)
		if !reflect.DeepEqual(got, n.want) {
			t.Errorf("%v = %v, want %v", n.num, got, n.want)
		}
	}
}
