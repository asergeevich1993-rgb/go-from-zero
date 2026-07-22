package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/repeat", func(w http.ResponseWriter, r *http.Request) {

		word := r.URL.Query().Get("word")
		strcount := r.URL.Query().Get("count")
		count, _ := strconv.Atoi(strcount)

		for i := 0; i < count; i++ {
			fmt.Fprint(w, word)
		}

	})
	http.ListenAndServe(":8080", nil)

}
