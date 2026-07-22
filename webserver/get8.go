package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/reverse", func(w http.ResponseWriter, r *http.Request) {

		word := r.URL.Query().Get("word")
		letters := []rune(word)
		rev := ""
		for i := len(letters) - 1; i >= 0; i-- {
			rev += string(letters[i])
		}
		fmt.Fprint(w, rev)

	})
	http.ListenAndServe(":8080", nil)
}
