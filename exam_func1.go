package main

import "fmt"

func sumSlice1(nums []int) int {
	sum := 0
	for _, num := range nums {

		sum = sum + num
	}
	return sum
}
func main() {
	result := sumSlice1([]int{1, 2, 3, 4, 5, 6})
	fmt.Print(result)

}
