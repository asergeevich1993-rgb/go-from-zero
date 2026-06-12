package main

import "fmt"

func main() {
	var n int
	sum := 0

	fmt.Println("Введите N: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {

		if i%2 != 0 {
			sum = sum + i
		}
	}
	fmt.Println("Сумма нечетных: ", sum)

}
