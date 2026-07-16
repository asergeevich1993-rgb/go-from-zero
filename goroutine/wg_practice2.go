package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 3)
	var wg sync.WaitGroup
	sum := 0
	nums := []int{100, 200, 300}

	for _, num := range nums {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			ch <- n

		}(num)
	}
	wg.Wait()
	close(ch)
	for n := range ch {

		fmt.Println(n)
		sum = sum + n
		fmt.Println(sum)
	}

}
