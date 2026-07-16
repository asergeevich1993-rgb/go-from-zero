package main

import (
	"fmt"
	"sync"
)

func main() {
	result := make(chan int, 10)
	tasks := make(chan int, 10)

	var wg sync.WaitGroup

	for w := 1; w <= 4; w++ {
		wg.Add(1)
		go func(WorkerID int) {
			defer wg.Done()
			for task := range tasks {
				results := task + 10
				result <- results
				fmt.Printf("Workers %d: %d + 10 = %d\n", WorkerID, task, results)
			}

		}(w)
	}

	for i := 1; i <= 8; i++ {
		tasks <- i
	}
	close(tasks)
	wg.Wait()
	close(result)
	sum := 0
	for r := range result {
		sum += r

	}
	fmt.Println(sum)
}
