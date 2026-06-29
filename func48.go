package main

import (
	"fmt"
	"strconv"
)

func countDigit(nums []int) map[int]int {

	result := map[int]int{}

	for _, num := range nums {
		length := len(strconv.Itoa(num))
		result[length] = result[length] + 1

	}
	return result
}
func main() {

	result := countDigit([]int{5, 12, 345, 67, 8})
	fmt.Print(result)

}
