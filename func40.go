package main

import (
	"fmt"
	"strings"
)

func groupByVowels(words []string) map[int][]string {
	result := map[int][]string{}
	for _, word := range words {
		count := 0
		letters := []rune(word)
		for _, letter := range letters {
			if strings.Contains("aeёиоэуыюя", string(letter)) {
				count++

			}
		}
		result[count] = append(result[count], word)
	}
	return result
}

func main() {

	result := groupByVowels([]string{"го", "это", "круто", "я", "ты"})
	fmt.Print(result)

}
