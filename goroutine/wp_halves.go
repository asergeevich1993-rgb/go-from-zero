package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	nums := []int{10, 20, 30, 40, 50, 60}
	results := make(chan int, 10)
	tasks := make(chan int, 10)
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(WorkerID int) {
			defer wg.Done()
			for task := range tasks {
				result := task / 2
				results <- result
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
