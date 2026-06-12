package main

import "fmt"

func main() {

	sum := 0
	var num int
	var count int

	fmt.Print("Сколько чисел: ")
	fmt.Scan(&count)

	for i := 1; i <= count; i++ {
		fmt.Println("Введите число: ")
		fmt.Scan(&num)

		sum = sum + num

	}

	fmt.Println("Сумма: ", sum)
	stats := float64(sum) / float64(count)
	fmt.Printf("Среднее: %.2f\n", stats)

}
