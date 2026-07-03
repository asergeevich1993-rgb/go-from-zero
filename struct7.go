package main

import "fmt"

type Product struct {
	Name string

	Price int
}

func groupByPrice(products []Product) map[int][]Product {
	result := map[int][]Product{}

	for _, p := range products {
		result[p.Price] = append(result[p.Price], p)
	}
	return result
}

func main() {
	products := []Product{
		{Name: "ручки", Price: 80},
		{Name: "тетради", Price: 120},
		{Name: "карандаши", Price: 100},
		{Name: "ластики", Price: 150}}

	result := groupByPrice(products)
	for key, name := range result {
		fmt.Print(key, ": ")
		for _, n := range name {
			fmt.Print(n.Name, " ")
		}
		fmt.Println()
	}

}
