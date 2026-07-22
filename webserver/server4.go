package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		first := r.URL.Query().Get("first")
		if first == "" {
			first = "гость"
		}
		last := r.URL.Query().Get("last")
		if last == "" {
			last = "гость"
		}

		fmt.Fprintf(w, "Привет, %s %s !", first, last)
	})
	http.ListenAndServe(":8080", nil)

}
