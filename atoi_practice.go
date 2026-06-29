package main

import (
	"fmt"
	"strconv"
)

func stringnum(text string) (int, bool) {

	num, err := strconv.Atoi(text)
	if err == nil {
		return num, true
	} else {
		return num, false
	}

}
func main() {

	result, num := stringnum("abc")
	fmt.Print(result, " ", num)
}
