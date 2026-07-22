package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "гость"
		}
		fmt.Fprintf(w, "привет, %s", name)
	})
	http.ListenAndServe(":8080", nil)

}
