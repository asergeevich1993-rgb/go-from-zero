package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/max", func(w http.ResponseWriter, r *http.Request) {

		max := 0
		astr := r.URL.Query().Get("a")
		bstr := r.URL.Query().Get("b")

		a, _ := strconv.Atoi(astr)
		b, _ := strconv.Atoi(bstr)

		if a > b {
			max = a
			fmt.Fprintf(w, "Максимум: %d", max)
		} else {
			max = b
			fmt.Fprintf(w, "Максимум: %d", max)
		}

	})
	http.ListenAndServe(":8080", nil)

}
