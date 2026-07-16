package main

import "fmt"

func main() {

	gen := func() <-chan int {
		out := make(chan int)
		go func() {
			for i := 1; i <= 5; i++ {
				out <- i
			}
			close(out)
		}()
		return out
	}
	mult := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n * 2
			}
			close(out)

		}()
		return out
	}
	numbers := gen()
	doubled := mult(numbers)
	for result := range doubled {
		fmt.Println(result)
	}
}
