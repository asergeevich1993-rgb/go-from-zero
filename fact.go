package main

import "fmt"

func main() {

	var N int
	fact := 1

	fmt.Print("введите N: ")
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {

		fact = fact * i

	}
	fmt.Print("Факториал: ", fact)
}
