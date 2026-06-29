package main

import (
	"fmt"
	"strings"
)

func chekletter(text string) bool {
	result := false
	if strings.Contains("оаэиеы", text) {
		result = true
		return result
	} else {
		return false
	}

}

func main() {
	result := chekletter("ы")
	fmt.Print(result)

}
