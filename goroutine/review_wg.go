package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	nums := []int{5, 10, 15, 20, 25, 30, 35, 40}
	tasks := make(chan int, 8)
	result := make(chan int, 8)
	for w := 1; w <= 4; w++ {
		wg.Add(1)
		go func(WorkerID int) {
			defer wg.Done()
			for task := range tasks {
				result <- task * 3
			}
		}(w)
	}

	for _, n := range nums {
		tasks <- n
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
