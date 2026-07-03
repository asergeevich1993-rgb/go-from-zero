package main

import "fmt"

type Person struct {
	Name string

	Age int

	Adult bool
}

func (user Person) Print1() {

	fmt.Println(user.Name, ":", user.Age)
}

func (user Person) isAdult() bool {
	return user.Age >= 18
}

func main() {
	p := Person{Name: "Артур", Age: 33}
	p.Print1()
	fmt.Print("Совершеннолетний: ", p.isAdult())
}
