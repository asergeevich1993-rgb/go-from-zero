package main

import "fmt"

func minmax(nums []int) (int, int) {
	min := nums[0]
	max := nums[0]

	for _, value := range nums {
		if min > value {
			min = value
		}
		if max < value {
			max = value
		}

	}
	return min, max

}

func main() {

	min, max := minmax([]int{3, 7, 1, 9, 4})

	fmt.Println(min, max)
}
