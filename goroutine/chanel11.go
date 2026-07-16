package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	sum := 0
	ch := make(chan int)

	go func() {

		for _, r := range slice {
			sum = sum + r
		}
		ch <- sum
	}()
	result := <-ch

	fmt.Println(result)
}
