package main

import "testing"

func maxInSlice(nums []int) int {
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

func TestMaxInSlice(t *testing.T) {
	num := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{10, 3, 4, 1}, 10},
		{[]int{-1, -2, 0, 1}, 1}}

	for _, n := range num {
		got := maxInSlice(n.nums)
		if got != n.want {
			t.Errorf("%v = %d, want %d", n.nums, got, n.want)
		}
	}
}
