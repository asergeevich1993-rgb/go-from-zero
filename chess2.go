package main

import "fmt"

func main() {
	var n int
	fmt.Print("высота: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("#")

			} else {
				fmt.Print("_")
			}
		}
		fmt.Println(" ")
	}

}
