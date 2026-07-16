package main

import "fmt"

func main() {

	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	nums1 := <-ch
	nums2 := <-ch
	nums3 := <-ch
	fmt.Println(nums1, nums2, nums3)
}
