package main

import "fmt"

func groupBylastLetter(words []string) map[string][]string {
	result := map[string][]string{}

	for _, word := range words {
		lastletter := string([]rune(word)[len([]rune(word))-1])
		result[lastletter] = append(result[lastletter], word)
	}
	return result
}
func main() {

	result := groupBylastLetter([]string{"арбуз", "ананас", "банан", "киви", "манго"})

	fmt.Print(result)

}
