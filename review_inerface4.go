package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Меня зовут " + p.Name

}

type Dog struct {
	Breed string
}

func (d Dog) Speak() string {
	return "Гав гав я, " + d.Breed

}

type Robot struct {
	Model string
}

func (r Robot) Speak() string {
	return "пик пик я, " + r.Model

}

func Chorus(speakers []Speaker) {
	for _, s := range speakers {
		fmt.Println(s.Speak())
	}
}

func main() {
	dog := Dog{Breed: "Овчарка"}
	person := Person{Name: "Артур"}
	robot := Robot{Model: "R2D2"}

	result := []Speaker{dog, person, robot}

	Chorus(result)
}
