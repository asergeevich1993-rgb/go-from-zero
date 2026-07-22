package main

import "fmt"

func isPalindrome(word string) bool {

	letters := []rune(word)

	for i := 0; i < len(letters); i++ {
		if letters[i] != letters[len(letters)-i-1] {
			return false
		}

	}
	return true

}

func main() {

	text := "kazak"
	result := isPalindrome(text)
	fmt.Println(result)
}
