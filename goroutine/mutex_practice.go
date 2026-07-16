package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	slice := []int{}
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			slice = append(slice, 2)
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(slice)
	sum := 0
	for _, s := range slice {
		sum += s
	}
	fmt.Println(sum)

}
