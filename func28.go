package main

import "fmt"

func minPrice(m map[string]int) (string, int) {
	var cheapest string
	minprice := 9999

	for key, value := range m {
		if value < minprice {
			minprice = value
			cheapest = key
		}
	}
	return cheapest, minprice

}

func main() {

	cheapest, minprice := minPrice(map[string]int{"ручки": 100, "тетради": 50, "карандаши": 200})
	fmt.Print(cheapest, ":", minprice)

}
