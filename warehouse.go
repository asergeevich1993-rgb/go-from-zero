package main

import "fmt"

func main() {

	warehouse := map[string]int{"яблоки": 100, "бананы": 50, "груши": 75}

	var warename string
	fmt.Print("Что ищем? ")
	fmt.Scan(&warename)

	ware, exists := warehouse[warename]
	if exists {
		fmt.Print(warename, ":", ware)
	} else {
		fmt.Print("не найдено")
	}
}
