package main

import (
	"encoding/json"
	"net/http"
)

type Greet struct {
	Greet string
}

func main() {

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {

		name := r.URL.Query().Get("name")
		if name == "" {
			name = "гость"

		}

		g := Greet{Greet: "привет," + name + "!"}
		data, _ := json.Marshal(g)
		w.Write(data)
	})

	http.ListenAndServe(":8080", nil)
}
