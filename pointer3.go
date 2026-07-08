package main

import "fmt"

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp

}

func main() {
	x := 5
	y := 10
	swap(&x, &y)
	fmt.Println(x, " ", y)
}
