package main

import "fmt"

func main() {

	var n int
	countdiv := 0
	sumdiv := 0

	fmt.Print("Введите N: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			countdiv++
			sumdiv = sumdiv + i
		}

	}
	fmt.Println("Колличество делителей: ", countdiv)
	fmt.Println("Сумма делителей: ", sumdiv)

}
