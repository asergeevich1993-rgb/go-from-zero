package main

import "fmt"

func main() {

	var count int
	var num int
	negsum := 0

	fmt.Print("Сколько чисел? ")
	fmt.Scan(&count)

	for i := 1; i <= count; i++ {
		fmt.Println("Введите число: ")
		fmt.Scan(&num)

		if num < 0 {

			negsum++
		}

	}
	fmt.Println("отрицательных: ", negsum)

}
