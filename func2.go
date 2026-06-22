package main

import "fmt"

func multiply(a int, b int, c int) int {
	return a * b * c
}

func main() {
	result := multiply(2, 3, 4)
	fmt.Println("Сумма: ", result)
}
