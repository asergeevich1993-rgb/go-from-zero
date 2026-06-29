package main

import "fmt"

func fistlastletter(text string) (string, string) {
	fistletter := string([]rune(text)[0])

	lastletter := []rune(text)
	last := string((lastletter)[len(lastletter)-1])
	return fistletter, last
}

func main() {

	fist, last := fistlastletter("мандарин")
	fmt.Print(fist, " ", last)
}
