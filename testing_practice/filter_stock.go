package main

import (
	"reflect"
	"testing"
)

func filterByPrice(stock map[string]int, maxPrice int) (map[string]int, error) {

	newstock := map[string]int{}
	for key, value := range stock {
		if value <= maxPrice {
			newstock[key] = value

		}

	}
	return newstock, nil
}

func TestFilterByPrice(t *testing.T) {
	result := []struct {
		stock    map[string]int
		maxPrice int
		want     map[string]int
		Errwant  bool
	}{
		{map[string]int{"BMW": 2000, "Toyota": 1500, "Dodge": 2500}, 2000, map[string]int{"BMW": 2000, "Toyota": 1500}, false},
		{map[string]int{"Warcraft": 100, "Cyberpunk": 150}, 100, map[string]int{"Warcraft": 100}, false},
		{map[string]int{"Creatine": 50}, 40, map[string]int{}, true}}

	for _, r := range result {
		got, err := filterByPrice(r.stock, r.maxPrice)
		if r.Errwant {
			if err == nil {
				t.Errorf("Ожидалась ошибка")
			}
		} else {
			if !reflect.DeepEqual(got, r.want) {
				t.Errorf("%v = %v, want %v", r.stock, got, r.want)
			}
			if err != nil {
				t.Errorf("непредвиденная ошибка")
			}
		}
	}
}
