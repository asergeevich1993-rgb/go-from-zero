package main

import "fmt"

func groupByRemainder(nums []int, divisor int) map[int][]int {
	result := map[int][]int{}
	remainder := 1
	for _, num := range nums {
		remainder = num % divisor
		result[remainder] = append(result[remainder], num)
	}
	return result
}

func main() {
	result := groupByRemainder([]int{10, 15, 22, 33, 40, 55}, 10)
	fmt.Print(result)

}
