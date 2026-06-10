package main

import "fmt"

func main() {

	var num int
	var count int
	sum := 0

	fmt.Println("Сколько чисел: ")
	fmt.Scan(&count)

	for i := 1; i <= count; i++ {
		fmt.Println("Введите число: ")
		fmt.Scan(&num)
		sum = sum + num

	}
	average := float64(sum) / float64(count)

	fmt.Printf("Среднее: %.2f\n", average)

}
