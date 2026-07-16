package main

import (
	"fmt"
	"sync"
)

func main() {
	result := make(chan int, 10)
	tasks := make(chan int, 10)

	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(WorkerID int) {
			defer wg.Done()
			for task := range tasks {
				if task%2 == 0 {
					results := task
					result <- results
					fmt.Printf("worker %d:%d\n", WorkerID, results)
				}
			}
		}(w)

	}
	for i := 1; i <= 10; i++ {
		tasks <- i
	}
	close(tasks)
	wg.Wait()
	close(result)
	for r := range result {
		fmt.Println(r)
	}

}
