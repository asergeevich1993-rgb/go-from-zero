package main

import "fmt"

func main() {

	cities := []string{"москва", "питер", "сочи", "казань", "омск"}
	found := false
	var city string
	fmt.Print("Введите город? ")
	fmt.Scan(&city)

	for _, v := range cities {
		if city == v {
			found = true
		} else {

		}

	}
	if found {
		fmt.Print("Найден")
	} else {
		fmt.Print("не найден")
	}

}
