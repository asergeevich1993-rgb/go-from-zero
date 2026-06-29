package main

import (
	"fmt"
	"strconv"
)

func groupByfistDigit(nums []int) map[string][]int {
	result := map[string][]int{}

	for _, num := range nums {
		str := strconv.Itoa(num)
		fistdigit := string((str)[0])
		result[fistdigit] = append(result[fistdigit], num)
	}

	return result

}

func main() {

	result := groupByfistDigit([]int{123, 45, 678, 12, 98, 456})
	fmt.Print(result)
}
