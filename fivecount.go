package main

import "fmt"

func main() {
	var count int
	var n int
	sumfive := 0

	fmt.Print("Сколько чисел? ")
	fmt.Scan(&count)

	for i := 1; i <= count; i++ {

		fmt.Print("Введите число ", i, ":")
		fmt.Scan(&n)

		if n == 5 {
			sumfive++
			fmt.Println("нашли пятерку, сейчас sumfive=", sumfive)
		}
	}
	fmt.Print("Сумма пятерок: ", sumfive)
}
