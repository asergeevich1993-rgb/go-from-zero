package main

import (
	"reflect"
	"testing"
)

func remove(nums []int, target int) []int {
	slice := []int{}
	for _, n := range nums {
		if n == target {
			continue
		}
		slice = append(slice, n)

	}
	return slice
}

func TestRemove(t *testing.T) {
	result := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4}},
		{[]int{4, 1, 2, 7, 8}, 7, []int{4, 1, 2, 8}},
		{[]int{7, 9}, 9, []int{7}}}
	for _, r := range result {
		got := remove(r.nums, r.target)
		if !reflect.DeepEqual(got, r.want) {
			t.Errorf("%v = %v,want %v", r.nums, got, r.want)

		}

	}
}
