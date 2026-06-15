package main

import "fmt"

func main() {
	var n int
	fmt.Print("высота: ")
	fmt.Scan(&n)

	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println(" ")
	}

}
