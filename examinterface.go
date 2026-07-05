package main

import (
	"fmt"
	"strconv"
)

type Describer interface {
	Describe() string
}
type Product struct {
	Name  string
	Price int
}

func (p Product) Describe() string {

	return p.Name + ": " + strconv.Itoa(p.Price)

}

type City struct {
	Name       string
	Population int
}

func (c City) Describe() string {
	return c.Name + " :" + strconv.Itoa(c.Population)

}

func main() {
	result := []Describer{Product{Name: "огурец", Price: 100}, City{Name: "Питер", Population: 200600}}
	for _, r := range result {
		fmt.Println(r.Describe())
	}
}
