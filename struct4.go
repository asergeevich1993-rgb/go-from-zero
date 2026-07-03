package main

import "fmt"

type Product struct {
	name  string
	price int
}

func cheapProduct(products []Product, maxprice int) []Product {
	stock := []Product{}

	for _, p := range products {
		if p.price <= maxprice {
			stock = append(stock, p)
		}

	}
	return stock

}
func main() {

	Products := []Product{
		{name: "ручки", price: 80},
		{name: "тетради", price: 120},
		{name: "карандаши", price: 100},
		{name: "ластики", price: 150},
	}
	stock := cheapProduct((Products), 100)

	for _, s := range stock {
		fmt.Println(s.name, "-", s.price)
	}
}
