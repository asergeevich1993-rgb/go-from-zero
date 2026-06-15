package main

import "fmt"

func main() {
	var N int

	fmt.Print("размер: ")
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("#")
			} else {
				fmt.Print("_")
			}
		}
		fmt.Println("")
	}
}
