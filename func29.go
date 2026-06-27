package main

import "fmt"

func sumSlice(nums []int) int {
	sum := 0
	for _, value := range nums {
		sum = sum + value
	}
	return sum
}

func main() {

	result := sumSlice([]int{10, 20, 30})
	fmt.Print(result)
}
