package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		city := r.URL.Query().Get("city")
		if city == "" {
			city = "гость"
			fmt.Fprintf(w, "Добро пожаловать %s", city)
			return
		}
		fmt.Fprintf(w, "Добро пожаловать в %s", city)
	})
	http.ListenAndServe(":8080", nil)

}
