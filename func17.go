package main

import "fmt"

func countTarget(nums []int, target int) int {
	count := 0
	for _, value := range nums {
		if value == target {
			count++
		}
	}
	return count
}

func main() {

	result := countTarget([]int{1, 2, 3, 2, 2, 4}, 2)
	fmt.Print("Количество: ", result)
}
