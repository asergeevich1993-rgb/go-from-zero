package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func mostExpensive(products []Product) Product {
	var mostexp Product
	fist := true
	for _, p := range products {
		if fist {
			mostexp = p
			fist = false
		}
		if p.Price > mostexp.Price {
			mostexp = p

		}
	}
	return mostexp

}
func main() {

	Products := []Product{
		{Name: "ручки", Price: 80},
		{Name: "тетради", Price: 120},
		{Name: "карандаши", Price: 100},
		{Name: "ластики", Price: 150},
	}
	stock := mostExpensive(Products)

	fmt.Println(stock.Name, ":", stock.Price)

	//fmt.Println(stock)
}
