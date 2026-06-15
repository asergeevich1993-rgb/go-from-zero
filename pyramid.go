package main

import (
	"fmt"
)

func main() {
	var N int

	fmt.Print("высота: ")
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		for j := 1; j <= N-i; j++ {
			fmt.Print("_")
		}
		for j := 1; j <= (2*i)-1; j++ {
			fmt.Print("#")
		}
		fmt.Println("")
	}

}
