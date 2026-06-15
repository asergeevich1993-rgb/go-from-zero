package main

import (
	"fmt"
)

func main() {
	var N int

	fmt.Print("высота: ")
	fmt.Scan(&N)
	mid := (N + 1) / 2
	for i := 1; i <= mid; i++ {
		for j := 1; j <= mid-i; j++ {
			fmt.Print("_")
		}
		for j := 1; j <= (2*i)-1; j++ {
			fmt.Print("#")
		}
		fmt.Println("")
	}

	for i := mid - 1; i >= 1; i-- {
		for j := 1; j <= mid-i; j++ {
			fmt.Print("_")
		}
		for j := 1; j <= (2*i)-1; j++ {
			fmt.Print("#")
		}
		fmt.Println(" ")
	}

}
