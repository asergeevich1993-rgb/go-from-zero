package main

import "fmt"

type Describer interface {
	Describe() string
}
type Vegetable struct {
	Name  string
	Color string
}

func (v Vegetable) Describe() string {
	return "Овощь: " + v.Name + " цвет: " + v.Color

}

type Fruit struct {
	Name  string
	Color string
}

func (f Fruit) Describe() string {
	return "Фрукт: " + f.Name + " цвет: " + f.Color

}

func show(d Describer) {
	fmt.Println(d.Describe())
}

func main() {
	fruit := Fruit{Name: "Яблоко", Color: "зеленое"}
	vegetable := Vegetable{Name: "огурчик", Color: "зеленый"}

	result := []Describer{fruit, vegetable}

	for _, r := range result {
		show(r)
	}

}
