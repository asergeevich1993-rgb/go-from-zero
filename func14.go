package main

import "fmt"

func sumPositive(nums []int) int {
	sum := 0

	for _, value := range nums {
		if value < 0 {
			continue
		}
		sum = value + sum
	}
	return sum
}

func main() {

	resuls := sumPositive([]int{3, -1, 4, -5, 2})
	fmt.Print("Сумма: ", resuls)

}
