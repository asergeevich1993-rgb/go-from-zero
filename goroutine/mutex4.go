package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	s := []int{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			s = append(s, 1)
			mu.Unlock()

		}()
	}
	wg.Wait()
	fmt.Println(s)

}
