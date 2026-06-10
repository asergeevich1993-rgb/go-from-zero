package main

import "fmt"

func main() {
	var number int
	fmt.Println("Введите число: ")
	fmt.Scan(&number)
	for i := 1; i <= number; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}

	fmt.Println("неужели я что то сделал сам")
}
