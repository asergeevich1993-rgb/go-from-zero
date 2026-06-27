package main

import "fmt"

func maxPrice(m map[string]int) (string, int) {
	var maxName string
	var maxPrice int

	for name, value := range m {
		if value > maxPrice {
			maxPrice = value
			maxName = name
		}
	}
	return maxName, maxPrice

}

func main() {

	prices := map[string]int{"ручки": 100, "тетради": 50, "карандаши": 200}

	maxName, maxPrice := maxPrice(prices)

	fmt.Print("макс: ", maxName, ":", maxPrice)

}
