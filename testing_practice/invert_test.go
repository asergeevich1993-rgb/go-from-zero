package main

import (
	"reflect"
	"testing"
)

func invert(m map[string]int) map[int]string {

	invertmap := map[int]string{}
	for key, value := range m {
		invertmap[value] = key
	}
	return invertmap
}

func TestInvert(t *testing.T) {
	result := []struct {
		m    map[string]int
		want map[int]string
	}{
		{map[string]int{"Ручки": 100, "Ластики": 50}, map[int]string{100: "Ручки", 50: "Ластики"}},
		{map[string]int{"BMW": 2020, "Toyota": 2021, "Nissan": 2024}, map[int]string{2020: "BMW", 2021: "Toyota", 2024: "Nissan"}}}

	for _, r := range result {
		got := invert(r.m)
		if !reflect.DeepEqual(got, r.want) {
			t.Errorf("%v = %v, want %v", r.m, got, r.want)
		}
	}
}
