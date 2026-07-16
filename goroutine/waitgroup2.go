package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	wg.Add(2)

	go func() {
		defer wg.Done()
		ch1 <- 10
	}()

	go func() {
		defer wg.Done()
		ch2 <- 20
	}()

	wg.Wait()
	num1 := <-ch1
	num2 := <-ch2

	fmt.Println(num1, " ", num2)
}
