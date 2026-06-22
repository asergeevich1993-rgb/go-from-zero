package main

import "fmt"

func main() {

	prices := map[string]int{"ручки": 100, "тетради": 50, "карандаши": 200, "ластики": 30}
	minPrice := 999999
	var cheapest string

	for name, value := range prices {
		if value < minPrice {
			minPrice = value
			cheapest = name
		}

	}
	fmt.Println("Самый дешевый: ", cheapest, ":", minPrice)

}
