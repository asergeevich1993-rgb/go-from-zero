package main

import "fmt"

func groupByLetterAny(words []string, letter string) map[int][]string {

	result := map[int][]string{}

	for _, word := range words {
		count := 0
		letters := []rune(word)
		for _, lttr := range letters {
			if letter == string(lttr) {
				count++
			}
		}
		result[count] = append(result[count], word)
	}
	return result
}

func main() {
	result := groupByLetterAny([]string{"арбуз", "банан", "апельсин", "киви"}, "и")
	fmt.Print(result)

}
