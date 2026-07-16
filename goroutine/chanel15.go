package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 5
		close(ch)

	}()

	for nums := range ch {
		fmt.Println(nums)
	}
}
