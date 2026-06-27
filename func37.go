package main

import "fmt"

func groupBySign(nums []int) map[string][]int {
	result := map[string][]int{}
	for _, num := range nums {

		if num > 0 {
			result["positive"] = append(result["positive"], num)
		}
		if num < 0 {
			result["negative"] = append(result["negative"], num)
		}
		if num == 0 {
			result["zero"] = append(result["zero"], num)
		}

	}
	return result

}
func main() {

	result := groupBySign([]int{-3, 0, 5, -1, 0, 8})

	fmt.Print(result)

}
