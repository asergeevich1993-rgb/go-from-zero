package main

import "fmt"

func main() {

	nums := []int{-1, 2, 4, 10, 100, 1001}
	max := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}

	}
	fmt.Print("максимум: ", max)

}
