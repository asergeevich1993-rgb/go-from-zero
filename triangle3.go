package main

import "fmt"

func main() {
	var height int

	fmt.Print("Высота: ")
	fmt.Scan(&height)

	for i := 1; i <= height; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
