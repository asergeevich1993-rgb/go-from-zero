package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Book struct {
	Title  string
	Author string
}

var books = []Book{
	{Title: "Hobbit", Author: "Tolkin"},
	{Title: "Go", Author: "Atrur"}}

func main() {

	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "PUT" {

			parts := strings.Split(r.URL.Path, "/")
			index, err := strconv.Atoi(parts[2])
			if err != nil {
				fmt.Fprintf(w, "Ошибка %v ", err)
				return
			}
			body, _ := io.ReadAll(r.Body)
			defer r.Body.Close()

			var newbook Book
			json.Unmarshal(body, &newbook)
			books[index] = newbook

			fmt.Fprint(w, "Изменено")

		}
		if r.Method == "GET" {
			data, _ := json.Marshal(books)

			w.Write(data)
		}

	})
	http.ListenAndServe(":8080", nil)
}
