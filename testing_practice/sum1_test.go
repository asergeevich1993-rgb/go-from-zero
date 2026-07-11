package main

import "testing"

func sumSlice1(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum = sum + n
	}
	return sum

}

func TestSumSlice1(t *testing.T) {
	sum := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{4, 5, 1}, 10},
		{[]int{-1, 2, 0, -1}, 0},
	}
	for _, s := range sum {
		got := sumSlice1(s.nums)
		if got != s.want {
			t.Errorf("Sum slice %v=%v,want %d", s.nums, got, s.want)
		}
	}
}
