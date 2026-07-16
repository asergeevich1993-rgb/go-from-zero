package main

import "fmt"

type Shape interface {
	Area() float64
	Name() string
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	area := r.Width * r.Height
	return area
}

func (r Rectangle) Name() string {
	return "Прямоугольник: "
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	area := 3.14 * c.Radius * c.Radius
	return area
}
func (c Circle) Name() string {
	return "Круг: "
}

func main() {
	shape := []Shape{Rectangle{Width: 5, Height: 3}, Circle{Radius: 5}}
	for _, s := range shape {
		fmt.Println(s.Name(), s.Area())

	}
}
