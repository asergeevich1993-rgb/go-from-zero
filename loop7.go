package main

import "fmt"

func main() {
	var n int
	count := 0
	fmt.Println("Введите число: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			count++
			fmt.Println(count)
		}
	}
}
