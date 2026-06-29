package main

import (
	"fmt"
	"strings"
)

func groupByMultipleLetters(words []string, letters string) map[int][]string {

	result := map[int][]string{}

	for _, word := range words {
		count := 0
		lttrs := []rune(word)
		for _, lttr := range lttrs {
			if strings.Contains(letters, string(lttr)) {
				count++
			}
		}
		result[count] = append(result[count], word)
	}
	return result
}

func main() {
	resilt := groupByMultipleLetters([]string{"арбуз", "банан", "апельсин", "киви"}, "ае")
	fmt.Print(resilt)
}
