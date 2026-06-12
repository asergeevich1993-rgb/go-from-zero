package main

import "fmt"

func main() {
	var n int

	fmt.Print("Введите N: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {

		fmt.Printf("%d^3=%d\n", i, i*i*i)

	}

}
