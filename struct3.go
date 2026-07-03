package main

import "fmt"

// Шаг 1: объявляем тип Person
type Person struct {
	Name string
	Age  int
	City string
}

// Шаг 2: функция-фильтр
func filterAdults(people []Person) []Person {
	result := []Person{}       // пустой слайс для результата
	for _, p := range people { // проходим по всем людям
		if p.Age >= 18 { // если возраст >= 18
			result = append(result, p) // добавляем в результат
		}
	}
	return result
}

// Шаг 3: main
func main() {
	people := []Person{
		{Name: "Артур", Age: 33, City: "Москва"},
		{Name: "Мария", Age: 16, City: "Питер"},
		{Name: "Иван", Age: 41, City: "Казань"},
	}

	adults := filterAdults(people)

	for _, p := range adults {
		fmt.Println(p.Name, "-", p.Age)
	}
}
