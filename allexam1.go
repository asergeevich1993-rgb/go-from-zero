package main

import "fmt"

func evenFilter(nums []int) []int {
	result := []int{}

	for _, num := range nums {
		if num%2 == 0 {
			result = append(result, num)
		}

	}
	return result
}

func main() {
	result := evenFilter([]int{1, 2, 3, 4, 5, 6})
	fmt.Print(result)

}
