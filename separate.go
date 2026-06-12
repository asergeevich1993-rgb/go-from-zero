package main

import "fmt"

func main() {

	var num int
	var count int
	evensum := 0
	oddsum := 0

	fmt.Print("Сколько чисел?")
	fmt.Scan(&count)
	for i := 1; i <= count; i++ {
		fmt.Println("Введите число ", i, ": ")
		fmt.Scan(&num)

		if num%2 == 0 {
			evensum = evensum + num
		} else {
			oddsum = oddsum + num
		}
	}
	fmt.Print("сумма четных: ", evensum)
	fmt.Println("")
	fmt.Print("сумма нечетных: ", oddsum)
}
