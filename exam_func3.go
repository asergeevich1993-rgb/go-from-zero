package main

import (
	"fmt"
	"strings"
)

func countWords(text string) map[string]int {
	counter := map[string]int{}
	words := strings.Split(text, " ")
	for _, word := range words {
		counter[word]++
	}
	return counter

}

func main() {

	result := countWords("го го это круто го это я")
	fmt.Print(result)
}
