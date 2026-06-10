package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	sum := 0
	for i := 1; i <= n; i++ {
		sum = sum + i

	}
	fmt.Print(" Сумма от 1 до ", n, "=", sum)
}
