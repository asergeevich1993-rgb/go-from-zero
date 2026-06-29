package main

import "fmt"

func fistandLast(word string) (string, string) {

	fist := string([]rune(word)[0])
	last := string([]rune(word)[len([]rune(word))-1])

	return fist, last
}
func main() {
	fist, last := fistandLast("арбуз")

	fmt.Print(fist, last)

}
