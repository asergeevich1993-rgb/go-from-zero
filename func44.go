package main

import "fmt"

func filterPositive(nums []int) []int {

	result := []int{}
	for _, value := range nums {
		if value > 0 {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	result := filterPositive([]int{-3, 5, -1, 8, 0})
	fmt.Print(result)
}
