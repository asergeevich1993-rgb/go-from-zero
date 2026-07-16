package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	tasks := make(chan int, 10)
	result := []int{}

	for w := 1; w <= 5; w++ {
		wg.Add(1)
		go func(WorkerID int) {
			defer wg.Done()
			for task := range tasks {
				mu.Lock()
				result = append(result, task)
				mu.Unlock()

			}
		}(w)
	}
	for i := 1; i <= 10; i++ {
		tasks <- i
	}
	close(tasks)
	wg.Wait()
	fmt.Println(result)
}
