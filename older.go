package main

import "fmt"

func main() {

	var age1 int
	var age2 int

	fmt.Println("Введите возраст: ")
	fmt.Scan(&age1)
	fmt.Scan(&age2)

	if age1 < age2 {
		fmt.Println("Второй старше")

	} else if age1 > age2 {
		fmt.Println("Первый старше")
	} else {
		fmt.Println("Одного возраста")
	}

}
