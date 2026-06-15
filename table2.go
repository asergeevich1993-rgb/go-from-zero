package main

import "fmt"

func main() {
	var n int
	fmt.Print("высота: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			{
				fmt.Print(i*j, "\t")

			}

		}
		fmt.Println(" ")
	}
}
