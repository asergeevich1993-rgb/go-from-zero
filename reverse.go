package main

import "fmt"

func main() {

	var n int

	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	for i := n; i >= 1; i-- {
		fmt.Println(i)
	}
}
