package main

import (
	"fmt"
	"strings"
)

func wordCount(text string) map[string]int {
	count := map[string]int{}
	words := strings.Split(text, " ")
	for _, v := range words {
		count[v]++
	}
	return count
}

func main() {
	result := wordCount("го го это круто го это")

	fmt.Println("Число: ", result)
}
