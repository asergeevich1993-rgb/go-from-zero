package main

import "fmt"

func contain(nums []int, target int) bool {
	for _, value := range nums {
		if value == target {
			return true
		}

	}
	return false

}

func main() {

	result1 := contain([]int{1, 2, 3, 4}, 3)

	fmt.Println("Верно: ", result1)

	result2 := contain([]int{1, 2, 3, 4}, 5)

	fmt.Println("неверно: ", result2)

}
