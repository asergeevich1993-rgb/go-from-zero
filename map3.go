package main

import "fmt"

func main() {

	number := map[string]int{"Артур": 123456, "Мария": 654321}
	var name string

	fmt.Println("имя: ")
	fmt.Scan(&name)

	numname, exists := number[name]

	if exists {
		fmt.Print(numname)
	} else {
		fmt.Print("не найден")
	}

}
