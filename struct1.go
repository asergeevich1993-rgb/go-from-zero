package main

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
}

func main() {

	p := Person{Name: "Артур", Age: 33, City: "Москва"}
	fmt.Println("Имя:", p.Name)
	fmt.Println("Возраст:", p.Age)
	fmt.Println("Город:", p.City)

}
