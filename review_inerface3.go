package main

import (
	"fmt"
	"strconv"
)

type Describer interface {
	Describe() string
}

type City struct {
	Name       string
	Population int
}

func (c City) Describe() string {
	return c.Name + " " + strconv.Itoa(c.Population)
}

type Product struct {
	Name  string
	Price int
}

func (p Product) Describe() string {
	return p.Name + " " + strconv.Itoa(p.Price)

}

func main() {

	result := []Describer{Product{Name: "ручка", Price: 100}, City{Name: "Москва", Population: 100500}}

	for _, r := range result {
		fmt.Println(r.Describe())
	}

}
