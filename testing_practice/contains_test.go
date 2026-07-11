package main

import "testing"

func contains(nums []int, target int) bool {
	for _, n := range nums {
		if target == n {
			return true
		}
	}
	return false
}

func TestContains(t *testing.T) {
	test := []struct {
		nums   []int
		target int
		want   bool
	}{
		{[]int{4, 5, 2, 6, 1}, 1, true},
		{[]int{4, 1, 5, 3, 2}, 4, true},
		{[]int{4, 1, 5, 5, 6}, 0, false}}

	for _, tt := range test {
		got := contains(tt.nums, tt.target)

		if got != tt.want {
			t.Errorf("Contains(%d,%d)=%v,want %v", tt.nums, tt.target, got, tt.want)

		}
	}

}
