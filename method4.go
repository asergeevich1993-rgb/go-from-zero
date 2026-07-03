package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (u Person) Print() {
	fmt.Println("Имя: ", u.Name, "Возраст: ", u.Age)

}

func (b *Person) Birthday() {
	b.Age += 1
}

func main() {
	p := Person{Name: "Артур", Age: 33}
	p.Print()
	p.Birthday()
	p.Print()
}
