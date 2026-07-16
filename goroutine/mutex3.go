package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	s := []int{}
	ch := make(chan int, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			ch <- rand.Intn(100) + 1

		}()
	}
	wg.Wait()
	close(ch)
	for nums := range ch {
		s = append(s, nums)
	}
	fmt.Println(s)
}
