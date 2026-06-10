package main

import "fmt"

func main() {
	var count int
	var num int
	sum := 0

	fmt.Print("Сколько чисел? ")
	fmt.Scan(&count)

	for i := 1; i <= count; i++ {
		fmt.Printf("Введите число %d: ", i) // %d подставит i
		fmt.Scan(&num)
		sum = sum + num
	}

	fmt.Println("Сумма:", sum)
}
