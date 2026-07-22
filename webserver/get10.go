package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/palindrome", func(w http.ResponseWriter, r *http.Request) {
		word := r.URL.Query().Get("word")

		letter := []rune(word)

		for i := 0; i < len(letter)/2; i++ {
			if letter[i] != letter[len(letter)-1-i] {
				fmt.Fprint(w, "не палиндром")
				return

			}

		}
		fmt.Fprint(w, "палиндром")

	})
	http.ListenAndServe(":8080", nil)

}
