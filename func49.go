package main

import (
	"fmt"
	"strconv"
)

func lennums(nums int) int {

	str := strconv.Itoa(nums)
	length := len(str)
	return length

}

func main() {

	result := lennums(123)
	fmt.Print(result)
}
