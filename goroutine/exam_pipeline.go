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
	double := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n * 2
			}
			close(out)

		}()
		return out

	}
	sumOne := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n - 1

			}
			close(out)

		}()
		return out

	}
	nums := gen()
	multi := double(nums)
	result := sumOne(multi)

	for r := range result {
		fmt.Println(r)
	}

}
