package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 4)
	fmt.Println("Cтарт")
	go func() {
		ch <- 1
		ch <- 2

	}()
	go func() {
		ch <- 3
		ch <- 4
	}()

	//time.Sleep(5 * time.Second)
	close(ch)

	for nums := range ch {
		fmt.Println(nums)
	}
	fmt.Println("Финиш")
}
