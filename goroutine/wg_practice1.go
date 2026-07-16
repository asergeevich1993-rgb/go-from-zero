package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 5)

	wg.Add(5)

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

	go func() {
		defer wg.Done()
		ch <- 4
	}()
	go func() {
		defer wg.Done()
		ch <- 5
	}()

	wg.Wait()

	close(ch)
	for nums := range ch {
		fmt.Println(nums)
	}
}
