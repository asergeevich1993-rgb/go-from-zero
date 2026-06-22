package main

import "fmt"

func main() {

	prices := map[string]int{"ручки": 100, "тетради": 50, "карандаши": 200, "ластики": 30}
	minPrice := 99999
	maxPrice := 0
	var cheapest string
	var expensive string
	for name, price := range prices {
		if price > maxPrice {
			maxPrice = price
			expensive = name
		} else if price < minPrice {
			minPrice = price
			cheapest = name
		}

	}
	fmt.Println("Самый дешевый товар: ", cheapest, ":", minPrice)
	fmt.Println("Самый дорогой товар: ", expensive, ":", maxPrice)

}
