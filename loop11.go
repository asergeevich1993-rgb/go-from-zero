package main

import "fmt"

func main() {

	var num int
	var count int
	max := 0

	fmt.Println("Сколько чисел: ")
	fmt.Scan(&count)
	for i := 1; i <= count; i++ {
		fmt.Println("Введите число: ")
		fmt.Scan(&num)
		if num > max {
			max = num
		}

	}
	fmt.Print("Максимум: ", max)
}
