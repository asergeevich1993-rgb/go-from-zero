package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 10
	}()
	go func() {
		ch <- 20
	}()

	num1 := <-ch
	num2 := <-ch
	fmt.Print(num1, num2)
}
