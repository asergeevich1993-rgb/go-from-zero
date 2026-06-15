package main

import "fmt"

func main() {
	var height int

	fmt.Print("Высота: ")
	fmt.Scan(&height)

	for i := height; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println("")
	}
}
