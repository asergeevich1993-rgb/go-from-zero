package main

import "fmt"

func groupByDivisible(nums []int, d int) map[string][]int {

	result := map[string][]int{}
	for _, num := range nums {
		if num%d == 0 {
			result["divisible"] = append(result["divisible"], num)
		} else {
			result["notDivisible"] = append(result["notDivisible"], num)
		}
	}
	return result
}

func main() {
	result := groupByDivisible([]int{10, 15, 20, 25, 30}, 10)
	fmt.Print(result)

}
