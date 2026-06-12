package main

import "fmt"

func main() {

	var n int
	count := 0

	fmt.Print("Введите: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}

	}
	if count == 2 {
		fmt.Println("Простое")
	} else {
		fmt.Println("непростое")
	}

}
