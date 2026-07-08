package main

import "fmt"

func addOne(n *int) {
	*n = *n + 1

}

func main() {
	x := 5
	addOne(&x)
	fmt.Print(x)
}
