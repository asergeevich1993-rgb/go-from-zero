package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Book struct {
	Title  string
	Author string
}

var books []Book

func main() {

	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			body, _ := io.ReadAll(r.Body)

			defer r.Body.Close()
			var newbook Book

			json.Unmarshal(body, &newbook)
			books = append(books, newbook)
			fmt.Fprint(w, "добавлено")
		}
		if r.Method == "GET" {

			data, _ := json.Marshal(books)

			w.Write(data)
		}
	})
	http.ListenAndServe(":8080", nil)

}
