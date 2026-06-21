package main

import "fmt"

func main() {
	stock := map[string]int{"ручки": 100, "тетради": 50}

	for product, count := range stock {
		fmt.Println(product, ":", count)
	}
	stock["ручки"] = stock["ручки"] + 20
	stock["тетради"] = stock["тетради"] - 50
	for product, count := range stock {
		fmt.Println(product, ":", count)
	}

}
