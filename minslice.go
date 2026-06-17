package main

import "fmt"

func main() {

	nums := []int{20, 30, -1000, 345, 500}

	min := nums[0]

	for i := 0; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
		}

	}
	fmt.Print("минимум: ", min)

}
