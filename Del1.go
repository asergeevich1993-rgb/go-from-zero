package main

import "fmt"

func main() {
	var n int

	fmt.Println("Введите N: ")
	fmt.Scan(&n)
	fmt.Print("делители: ")
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
		}

	}
}
