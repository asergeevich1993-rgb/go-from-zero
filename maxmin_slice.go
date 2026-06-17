package main

import "fmt"

func main() {
	nums := []int{12, -5, 33, 0, 18}
	max := nums[0]
	min := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		} else if nums[i] < min {
			min = nums[i]

		}
		/*if nums[i] < min {
		min = nums[i]}
		*/

	}
	fmt.Print("максимум: ", max, "\n")
	fmt.Print("минимум: ", min)
}
