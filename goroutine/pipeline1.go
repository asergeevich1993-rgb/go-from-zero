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
	add := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n + 10

			}
			close(out)

		}()
		return out

	}
	nums := gen()
	nums2 := mult(nums)
	result := add(nums2)
	for r := range result {
		fmt.Println(r)
	}

}
