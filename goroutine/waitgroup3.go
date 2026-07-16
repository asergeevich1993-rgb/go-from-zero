package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 3)
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		ch <- 1

	}()
	go func() {
		defer wg.Done()
		ch <- 2
	}()
	go func() {
		defer wg.Done()
		ch <- 3

	}()
	wg.Wait()
	close(ch)

	for nums := range ch {
		fmt.Println(nums)
	}

}
