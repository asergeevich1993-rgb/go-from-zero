package main

import "fmt"

func main() {

	var n int
	fmt.Println("Введите число: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Printf("%d^2=%d\n", i, i*i)

	}
}
