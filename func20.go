package main

import "fmt"

func reverse(nums []int) []int {

	rev := []int{}

	for i := len(nums) - 1; i >= 0; i-- {
		rev = append(rev, nums[i])
	}

	return rev

}

func main() {

	result := reverse([]int{1, 2, 3, 4, 5})

	fmt.Print(result)

}
