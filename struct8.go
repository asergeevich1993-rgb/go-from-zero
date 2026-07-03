package main

import (
	"fmt"
)

type Product struct {
	Name  string
	Price int
}

func filterByPriceRange(products []Product, minprice, maxprice int) []Product {
	stock := []Product{}

	for _, p := range products {
		if p.Price >= minprice && p.Price <= maxprice {
			stock = append(stock, p)
		}
	}
	return stock

}

func main() {

	Products := []Product{
		{Name: "ручки", Price: 80},
		{Name: "тетради", Price: 120},
		{Name: "карандаши", Price: 100},
		{Name: "ластики", Price: 150}}

	result := filterByPriceRange((Products), 80, 120)
	fmt.Print(result)
	fmt.Println()
	for _, r := range result {
		fmt.Println(r)

	}

}
