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
	square := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n * n
			}
			close(out)
		}()
		return out
	}
	nums := gen()
	result := square(nums)

	for r := range result {
		fmt.Println(r)
	}
}
