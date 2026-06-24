package main

import "fmt"

func remove(nums []int, target int) []int {
	number := []int{}

	for _, value := range nums {
		if value == target {
			continue
		}
		number = append(number, value)
	}
	return number
}

func main() {

	result := remove([]int{1, 2, 3, 4, 5}, 2)

	fmt.Print(result)
}
