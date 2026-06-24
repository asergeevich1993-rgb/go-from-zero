package main

import "fmt"

func sumEven(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			sum = sum + nums[i]
		}

	}
	return sum

}

func main() {

	result := sumEven([]int{1, 2, 3, 4, 5, 6})
	fmt.Print("Сумма: ", result)

}
