package main

import "fmt"

func main() {

	age := map[string]int{"Артур": 33, "Мария": 25, "Иван": 41}

	for key, value := range age {
		fmt.Println(key, " :", value)
	}

}
