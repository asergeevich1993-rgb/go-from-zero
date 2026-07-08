package main

import "fmt"

func double(n *int) {
	*n = *n * 2
}

func main() {
	x := 5
	p := &x

	double(p)
	fmt.Print(*p)
}
