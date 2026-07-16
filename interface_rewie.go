package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "привет, я " + p.Name

}

type Robot struct {
	Name string
}

func (r Robot) Speak() string {
	return "привет, я " + r.Name

}
func Greet(s Speaker) {
	fmt.Println(s.Speak())

}
func main() {
	speakers := []Speaker{Person{Name: "Артур"}, Robot{Name: "R2D2"}}
	//for _, s := range speakers {
	//	fmt.Println(s.Speak())
	//}

	for _, s := range speakers {
		Greet(s)
	}
}
