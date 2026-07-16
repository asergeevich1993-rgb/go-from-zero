package main

import "fmt"

func main() {

	gen := func() <-chan int {
		out := make(chan int)
		go func() {
			for i := 1; i <= 10; i++ {
				out <- i
			}
			close(out)
		}()
		return out
	}
	filtereven := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				if n%2 == 0 {
					out <- n
				}
			}
			close(out)

		}()
		return out
	}

	multi := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n * 3
			}
			close(out)

		}()
		return out
	}

	nums := gen()
	result := filtereven(nums)
	done := multi(result)
	for r := range done {
		fmt.Println(r)
	}
}
