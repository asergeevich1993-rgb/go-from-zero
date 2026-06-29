package main

import "fmt"

func minMax(nums []int) (int, int) {
	min := nums[0]
	max := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}

	}
	return min, max
}

func main() {

	min, max := minMax([]int{7, 2, 9, 4})
	fmt.Print(min, max)
}
