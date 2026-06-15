package main

import "fmt"

func main() {

	var height int

	fmt.Print("высота: ")
	fmt.Scan(&height)

	for i := 1; i <= height; i++ { // строки
		for j := 1; j <= i; j++ { // звёздочки в строке
			fmt.Print("*")
		}
		fmt.Println() // переход на новую строку
	}
}
