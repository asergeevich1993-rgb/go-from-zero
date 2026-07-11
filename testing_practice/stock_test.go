package main

import (
	"fmt"
	"reflect"
	"testing"
)

func takeStock(stock map[string]int, item string, amount int) (map[string]int, error) {

	value, exists := stock[item]
	if exists && value >= amount {
		stock[item] = value - amount
		return stock, nil
	} else if exists && value < amount {
		return stock, fmt.Errorf("Не хватает на складе")
	} else {
		return stock, fmt.Errorf("Товара нет на складе")
	}

}

func TestTakeStock(t *testing.T) {
	result := []struct {
		m       map[string]int
		item    string
		amount  int
		want    map[string]int
		Errwant bool
	}{
		{map[string]int{"Ручки": 100, "Ластики": 50}, "Ластики", 40, map[string]int{"Ручки": 100, "Ластики": 10}, false},
		{map[string]int{"Тетради": 50}, "Тетради", 60, map[string]int{}, true},
		{map[string]int{"Линейки": 60, "Пеналы": 20}, "Ручки", 40, map[string]int{"Линейки": 60, "Пеналы": 20}, true},
	}
	for _, r := range result {
		got, err := takeStock(r.m, r.item, r.amount)
		if r.Errwant {
			if err == nil {
				t.Errorf("Ожидалась ошибка,но ее нет")
			}
		} else {
			if !reflect.DeepEqual(got, r.want) {
				t.Errorf("%v = %v, want %v", r.m, got, r.want)

			}
			if err != nil {
				t.Errorf("непредвиденная ошибка")
			}
		}

	}
}
