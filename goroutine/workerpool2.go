package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {

	tasks := make(chan string, 10)
	result := make(chan string, 10)
	words := []string{"hello", "world", "go", "is", "great", "fun"}
	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(WorkerID int) {
			defer wg.Done()
			for task := range tasks {
				results := strings.ToUpper(task)
				result <- results

			}
		}(w)

	}
	for _, word := range words {
		tasks <- word
	}
	close(tasks)
	wg.Wait()
	close(result)

	for r := range result {
		fmt.Println(r)
	}
}
