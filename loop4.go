package main

import "fmt"

func main() {
	var number int
	sum := 0
	fmt.Println("Введите число: ")
	fmt.Scan(&number)
	for i := 1; i <= number; i++ {
		sum = sum + i

	}
	fmt.Println("Сумма:", sum)

}
