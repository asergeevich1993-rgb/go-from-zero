package main

import (
	"fmt"
	"strings"
)

func main() {

	gen := func() <-chan string {
		words := []string{"го", "круто", "учим", "Go"}
		out := make(chan string)
		go func() {
			for _, word := range words {
				out <- word

			}
			close(out)

		}()
		return out
	}
	add := func(in <-chan string) <-chan string {
		out := make(chan string)
		go func() {
			for n := range in {
				out <- n + "!"
			}
			close(out)
		}()
		return out

	}
	toUpper := func(in <-chan string) <-chan string {
		out := make(chan string)
		go func() {
			for n := range in {
				result := strings.ToUpper(n)
				out <- result
			}
			close(out)

		}()
		return out
	}

	words := gen()
	result := add(words)
	done := toUpper(result)
	for r := range done {
		fmt.Println(r)
	}

}
