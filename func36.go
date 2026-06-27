package main

import (
	"fmt"
	"strconv"
)

func groupBylength(nums []int) map[int][]int {
	result := map[int][]int{}
	for _, num := range nums {
		length := len(strconv.Itoa(num))
		result[length] = append(result[length], num)
	}

	return result
}

func main() {
	result := groupBylength([]int{1, 23, 456, 78, 9})

	fmt.Print(result)

}
