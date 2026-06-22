package main

import "fmt"

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	result := isEven(3)
	fmt.Println("чет или нечет: ", result)
}
