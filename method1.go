package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Print() {

	fmt.Println(p.Name, "- ", p.Age)
}
func main() {
	p := Person{Name: "Артур", Age: 33}
	p.Print()

}
