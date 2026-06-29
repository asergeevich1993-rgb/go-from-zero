package main

import "fmt"

func filterPos(nums []int) []int {
	result := []int{}
	for _, num := range nums {
		if num > 0 {
			result = append(result, num)
		}
	}
	return result
}
func main() {

	result := filterPos([]int{-2, 2, -4, 5, 8, -10})
	fmt.Print(result)
}
