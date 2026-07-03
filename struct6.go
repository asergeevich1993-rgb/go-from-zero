package main

import "fmt"

type Product struct {
	Name string

	Price int
}

func AveragePrice(products []Product) float64 {
	var sum int
	var average float64
	for _, p := range products {
		sum = sum + p.Price

	}
	average = float64(sum) / float64(len(products))
	return average
}
func main() {
	products := []Product{
		{Name: "ручки", Price: 80},
		{Name: "тетради", Price: 120},
		{Name: "карандаши", Price: 100},
		{Name: "ластики", Price: 150}}

	average := AveragePrice(products)

	fmt.Print(average)

}
