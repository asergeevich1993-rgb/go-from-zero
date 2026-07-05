package main

import (
	"fmt"
	"strconv"
)

type Discraber interface {
	Discrabe() string
}

type Product struct {
	Name  string
	Price int
}

func (p Product) Discrabe() string {

	return "Товар: " + p.Name + ", цена: " + strconv.Itoa(p.Price)

}

type City struct {
	Name       string
	Population int
}

func (c City) Discrabe() string {
	return "Город: " + c.Name + ", Популяция: " + strconv.Itoa(c.Population)
}

func showInfo(d Discraber) {
	fmt.Println(d.Discrabe())

}

func main() {
	c := City{Name: "москва", Population: 13000000}
	p := Product{Name: "ручки", Price: 100}
	showInfo(c)
	showInfo(p)
}
