package main

import "testing"

func mininSlice(nums []int) int {
	min := nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}

func TestMininSlice(t *testing.T) {
	num := []struct {
		nums []int
		want int
	}{{[]int{2, 4, 5, 1}, 1},
		{[]int{0, -2, 5, 8}, -2},
		{[]int{7}, 7}}
	for _, n := range num {
		got := mininSlice(n.nums)
		if got != n.want {
			t.Errorf("%v=%v,want %d", n.nums, got, n.want)
		}
	}

}
