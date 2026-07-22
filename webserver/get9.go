package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/concat", func(w http.ResponseWriter, r *http.Request) {

		a := r.URL.Query().Get("a")
		b := r.URL.Query().Get("b")

		fmt.Fprint(w, a+b)

	})
	http.ListenAndServe(":8080", nil)

}
