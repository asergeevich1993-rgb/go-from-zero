package main

import "fmt"

func minMAX(nums []int) (int, int) {
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
	min, max := minMAX([]int{3, 4, 5, 10, 1, -1})
	fmt.Print(min, " ", max)
}
