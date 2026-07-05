package main

import "fmt"

// Интерфейс: любой, кто умеет говорить
type Speaker interface {
	Speak() string
}

// Человек
type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Привет, я " + p.Name
}

// Собака
type Dog struct {
	Breed string
}

func (d Dog) Speak() string {
	return "Гав! Я " + d.Breed
}

// Функция принимает любого Speaker'а
func greet(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	p := Person{Name: "Артур"}
	d := Dog{Breed: "овчарка"}

	greet(p) // Привет, я Артур
	greet(d) // Гав! Я овчарка
}
