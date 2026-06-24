package main

import "fmt"

func find(nums []int, target int) (int, bool) {
	for i, value := range nums {
		if value == target {
			return i, true
		}

	}
	return -1, false

}

func main() {

	index, found := find([]int{10, 20, 30, 40}, 30)

	fmt.Println(index, found)

	index1, found1 := find([]int{10, 20, 30, 40}, 50)

	fmt.Println(index1, found1)

}
