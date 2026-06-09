package main

import "fmt"

func main() {

	var color string

	fmt.Print("выведи цвет светофора (красный/желтый/зеленый): ")
	fmt.Scan(&color)

	if color == "красный" {
		fmt.Println("стой!")
	} else if color == "желтый" {
		fmt.Println("приготовся!")
	} else if color == "зеленый" {
		fmt.Println("иди!")
	} else {
		fmt.Println("неизвестный цвет!")
	}

}
