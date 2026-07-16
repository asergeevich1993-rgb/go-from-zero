package main

import "fmt"

func main() {

	gen := func() <-chan int {
		out := make(chan int)
		go func() {
			for i := 1; i <= 6; i++ {
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
	half := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n / 2
			}
			close(out)
		}()
		return out
	}

	nums := gen()
	sq := square(nums)
	result := half(sq)

	for r := range result {
		fmt.Println(r)
	}
}
