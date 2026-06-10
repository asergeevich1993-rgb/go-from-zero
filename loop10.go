package main

import "fmt"

func main() {

	var num int
	sum := 0
	count := 1

	fmt.Print("Сколько чисел: ")
	fmt.Scan(&count)
	for i := 1; i <= count; i++ {
		fmt.Println("Введите число: ")
		fmt.Scan(&num)

		if num%2 == 0 {
			sum = sum + num

		}

	}
	fmt.Print(sum)

}
