package main

import "fmt"

func main() {

	var n int

	fmt.Print("высота: ")
	fmt.Scan(&n)

	for i := n; i >= 1; i++ {
		for j := i; i >= j; j++ {
			fmt.Print(j)
		}
		fmt.Println("")
	}
}
