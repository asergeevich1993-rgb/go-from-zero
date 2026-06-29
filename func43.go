package main

import (
	"fmt"
	"strings"
)

func wordCount(text string) map[string]int {
	result := map[string]int{}
	words := strings.Split(text, " ")
	for _, word := range words {
		result[word]++
	}
	return result

}
func main() {
	result := wordCount("го го это круто го это")
	fmt.Print(result)
}
