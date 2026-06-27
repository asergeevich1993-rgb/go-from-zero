package main

import "fmt"

func groupByEvenOdd(nums []int) map[string][]int {
	result := map[string][]int{}
	for _, num := range nums {
		if num%2 == 0 {
			result["even"] = append(result["even"], num)
		} else {
			result["odd"] = append(result["odd"], num)
		}
	}
	return result
}

func main() {

	result := groupByEvenOdd([]int{1, 2, 3, 4, 5, 6})
	fmt.Print(result)

}
