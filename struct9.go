package main

import "fmt"

type Product struct {
	Name string

	Price int
}

func cheapest(products []Product) Product {

	var minprice Product
	fist := true

	for _, p := range products {
		if fist {
			minprice = p
			fist = false
		}
		if p.Price < minprice.Price {
			minprice = p

		}

	}
	return minprice

}

func main() {

	products := []Product{
		{Name: "ручки", Price: 80},
		{Name: "тетради", Price: 120},
		{Name: "карандаши", Price: 100},
		{Name: "ластики", Price: 150}}

	result := cheapest(products)

	fmt.Print(result)
}
