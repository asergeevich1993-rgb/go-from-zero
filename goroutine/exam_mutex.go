package main

import (
	"fmt"
	"sync"
)

func main() {

	nums := []int{}
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			nums = append(nums, 3)
			mu.Unlock()

		}()

	}
	wg.Wait()
	sum := 0
	for _, n := range nums {

		sum += n
	}
	fmt.Println(sum)

}
