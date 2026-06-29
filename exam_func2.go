package main

import "fmt"

func contains(nums []int, target int) bool {
	find := false

	for _, num := range nums {
		if num == target {
			find = true

		}
	}
	return find
}

func main() {
	result := contains([]int{1, 2, 3, 4, 5, 6}, 5)
	fmt.Print(result)
}
