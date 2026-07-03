package main

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	people := []Person{
		{Name: "Артур", Age: 33, City: "Москва"},
		{Name: "Мария", Age: 25, City: "Питер"},
		{Name: "Иван", Age: 41, City: "Казань"},
	}

	for _, p := range people {
		fmt.Println(p.Name, "-", p.Age, "лет,", p.City)
	}
}
