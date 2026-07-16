package main

import (
	"fmt"
	"sync"
)

func main() {

	var mu sync.Mutex
	var wg sync.WaitGroup
	var m map[string]int = map[string]int{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			mu.Lock()

			m["counter"]++
			mu.Unlock()
		}
	}()
	wg.Wait()
	fmt.Println(m)

}
