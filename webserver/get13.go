package main

import "fmt"

func reverse(s string) string {
	letter := []rune(s)
	reverse := ""
	for i := len(letter) - 1; i >= 0; i-- {
		reverse = reverse + string(letter[i])
	}
	return reverse

}

func main() {

	text := "Privet"
	result := reverse(text)

	fmt.Println(result)

}
