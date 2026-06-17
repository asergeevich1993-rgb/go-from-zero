package main

import "fmt"

func main() {

	nums := []int{-1, 2, 4, 10, 56, 1001}
	var average float64
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]

	}
	average = float64(sum) / float64(len(nums))

	fmt.Println("Сумма: ", sum)
	fmt.Println("Среднее: ", average)

}
