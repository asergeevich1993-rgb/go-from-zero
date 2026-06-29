package main

import (
	"fmt"
	"strconv"
)

func sumString(a, b string) int {

	sum1, err1 := strconv.Atoi(a)
	sum2, err2 := strconv.Atoi(b)

	if err1 != nil {
		sum1 = 0

	}
	if err2 != nil {
		sum2 = 0

	}
	return sum1 + sum2
}

func main() {

	result := sumString("xyz", "5")
	fmt.Print(result)

}
