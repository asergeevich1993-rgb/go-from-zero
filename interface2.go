package main

import "fmt"

type Printer interface {
	Print()
}

type Book struct {
	Name string
}

func (b Book) Print() {
	fmt.Println("Привет я, ", b.Name)
}

type Car struct {
	Name string
}

func (c Car) Print() {
	fmt.Println("Привет, а я, ", c.Name)

}

func ShowInfo(p Printer) {
	p.Print()
}
func main() {
	b := Book{Name: "Книга"}
	c := Car{Name: "Машина "}

	ShowInfo(b)
	ShowInfo(c)
}
