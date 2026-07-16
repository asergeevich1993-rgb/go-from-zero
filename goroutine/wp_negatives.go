package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	nums := []int{-5, 3, -2, 7, -1, 0, -8, 4}
	results := make(chan int, 10)
	tasks := make(chan int, 10)
	for w := 1; w <= 4; w++ {
		wg.Add(1)
		go func(WorkerID int) {
			defer wg.Done()
			for task := range tasks {
				if task < 0 {
					result := task
					results <- result
				}
			}

		}(w)

	}
	for _, num := range nums {
		tasks <- num

	}
	close(tasks)
	wg.Wait()
	close(results)
	for r := range results {
		fmt.Println(r)
	}

}
