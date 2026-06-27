package main

import "fmt"

func filter(m map[string]int, minPrice int) map[string]int {
	price := map[string]int{}
	for key, value := range m {
		if value >= minPrice {
			price[key] = value
		}
	}
	return price
}

func main() {

	price := filter(map[string]int{"ручки": 100, "тетради": 50, "карандаши": 200}, 100)

	fmt.Print(price)
}
