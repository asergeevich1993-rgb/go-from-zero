package main

import "fmt"

func main() {

	var number int

	fmt.Println("Введите число: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println("Четное")
	} else {
		fmt.Println("нечетное")
	}
}
