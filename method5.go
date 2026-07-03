package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func (prnt Product) Print() {
	fmt.Println("Товар:", prnt.Name, "Цена:", prnt.Price)
}

func (disc *Product) Discount(percent int) {
	disc.Price = disc.Price - disc.Price*percent/100
}

func main() {
	p := Product{Name: "ручки", Price: 100}
	p.Print()
	p.Discount(20)
	p.Print()
}
