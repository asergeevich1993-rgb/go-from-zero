package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func newPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

func main() {
	p := newPerson("Артур", 33)

	fmt.Println(p)
	fmt.Println(p.Name, ": ", p.Age)

}
