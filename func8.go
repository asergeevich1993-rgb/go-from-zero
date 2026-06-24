package main

import "fmt"

func minInSlice(nums []int) int {
	min := nums[0]
	for _, value := range nums {
		if min > value {
			min = value
		}

	}
	return min
}

func main() {

	result := minInSlice([]int{3, 7, 1, 9, 4})

	fmt.Print("Число: ", result)

}
