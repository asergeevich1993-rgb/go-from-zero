package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Первая")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Вторая")
	}()

	wg.Wait()
	fmt.Println("Финиш")
}
