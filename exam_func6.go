package main

import (
	"fmt"
	"strings"
)

func groupbyLetterwords(words []string, letter string) map[int][]string {
	result := map[int][]string{}
	for _, word := range words {
		count := 0
		letters := []rune(word)
		for _, letter := range letters {
			if strings.Contains("о", string(letter)) {
				count++

			}
		}
		result[count] = append(result[count], word)
	}
	return result
}
func main() {

	result := groupbyLetterwords([]string{"го", "это", "круто", "я", "ты", "огонь"}, "о")

	fmt.Print(result)
}
