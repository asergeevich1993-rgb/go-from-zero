package main

import "fmt"

func maxInSlice(nums []int) int {
	maxslice := 0
	for _, v := range nums {
		if v > maxslice {
			maxslice = v
		}

	}
	return maxslice
}

func main() {
	result := maxInSlice([]int{3, 7, 1, 9, 4})

	fmt.Println("MAX: ", result)
}
