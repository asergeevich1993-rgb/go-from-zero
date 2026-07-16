package main

import "fmt"

func main() {

	ch := make(chan string)

	go func() {
		ch <- "hello"
		ch <- "word"

	}()

	word1 := <-ch
	word2 := <-ch
	fmt.Println(word1, "", word2)
}
