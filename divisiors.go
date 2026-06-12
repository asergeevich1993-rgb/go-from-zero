package main

import "fmt"

func main() {
	var n int
	div := 0

	fmt.Println("Введите N:")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			div++
		}

	}
	fmt.Print("Колличество делителей: ", div)

}
