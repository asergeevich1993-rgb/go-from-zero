package main

import "fmt"

func average(nums []int) float64 {
	var sum int

	for _, value := range nums {
		sum = sum + value

	}

	return float64(sum) / float64(len(nums))

}

func main() {

	result := average([]int{10, 20, 30, 40})

	fmt.Println("Среднее: ", result)

}
