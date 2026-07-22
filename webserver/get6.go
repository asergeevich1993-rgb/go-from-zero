package main

import (
	"encoding/json"
	"net/http"
)

type Length struct {
	Word     string
	Lenwords int
}

func main() {

	http.HandleFunc("/length", func(w http.ResponseWriter, r *http.Request) {
		text := r.URL.Query().Get("text")

		t := Length{Word: text, Lenwords: len([]rune(text))}

		data, _ := json.Marshal(t)

		w.Write(data)
	})

	http.ListenAndServe(":8080", nil)
}
