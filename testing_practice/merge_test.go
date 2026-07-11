package main

import (
	"reflect"
	"testing"
)

func merge(a, b []int) []int {
	a = append(a, b...)
	return a
}

func TestMerge(t *testing.T) {
	result := []struct {
		a, b []int
		want []int
	}{
		{[]int{1, 2}, []int{3, 4}, []int{1, 2, 3, 4}},
		{[]int{5, 6, 7}, []int{8, 9, 10}, []int{5, 6, 7, 8, 9, 10}},
		{[]int{}, []int{1, 2, 3}, []int{1, 2, 3}}}
	for _, r := range result {
		got := merge(r.a, r.b)
		if !reflect.DeepEqual(got, r.want) {
			t.Errorf("%v + %v = %v, want %v", r.a, r.b, got, r.want)
		}
	}
}
