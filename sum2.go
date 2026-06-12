package main

import "fmt"

func main() {
	var n int
	sum := 0

	fmt.Print("Введите N: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		sum = sum + i*i
	}
	fmt.Print("сумма квадратов: ", sum)
}
