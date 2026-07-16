package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	tasks := make(chan int, 10)
	counter := map[string]int{}
	nums := []int{4, 8, 15, 16, 23, 42}

	for w := 1; w <= 4; w++ {
		wg.Add(1)
		go func(WorkerID int) {
			defer wg.Done()
			for range tasks {
				mu.Lock()
				counter["task"]++
				mu.Unlock()
			}
		}(w)
	}

	for _, n := range nums {
		tasks <- n
	}
	close(tasks)
	wg.Wait()
	fmt.Println(counter)

}
