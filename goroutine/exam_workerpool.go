package main

import (
	"fmt"
	"sync"
)

func main() {

	tasks := make(chan int, 10)
	results := make(chan int, 10)
	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for task := range tasks {
				result := task * 5
				results <- result
				fmt.Printf("worker %d: %d * 5 = %d\n", workerID, task, result)
			}
		}(w)
	}

	for i := 1; i <= 6; i++ {
		tasks <- i
	}
	close(tasks)
	wg.Wait()
	close(results)
	sum := 0
	for r := range results {
		sum += r

	}
	fmt.Println(sum)

}
