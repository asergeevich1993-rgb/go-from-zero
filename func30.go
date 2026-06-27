package main

import "fmt"

func find(nums []int, target int) (int, bool) {

	for _, value := range nums {
		if value == 10 {

			return value, true
		}
	}
	return -1, false

}
func main() {

	numbers, target := find([]int{5, 10, 15}, 10)

	fmt.Print(numbers, ":", target)

}
