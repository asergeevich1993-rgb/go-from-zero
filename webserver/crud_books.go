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

var books []Book

func main() {

	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Ошибка чтения %v", err)
				return
			}
			defer r.Body.Close()
			var b Book
			json.Unmarshal(body, &b)

			books = append(books, b)
			data, err := json.Marshal(books)
			if err != nil {
				fmt.Fprintf(w, "Ошибка сериализации %v", err)
				return
			}
			w.Write(data)

		}
		if r.Method == "GET" {
			data, _ := json.Marshal(books)
			w.Write(data)
		}

	})
	http.HandleFunc("/book/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "PUT" {
			parts := strings.Split(r.URL.Path, "/")
			index, _ := strconv.Atoi(parts[2])

			body, _ := io.ReadAll(r.Body)
			defer r.Body.Close()
			var b Book
			json.Unmarshal(body, &b)
			books[index] = b
			fmt.Fprint(w, "Обновлено")
		}
		if r.Method == "DELETE" {
			parts := strings.Split(r.URL.Path, "/")
			index, _ := strconv.Atoi(parts[2])

			books = append(books[:index], books[index+1:]...)
			fmt.Fprint(w, "удалено")
		}
		if r.Method == "GET" {
			parts := strings.Split(r.URL.Path, "/")
			index, _ := strconv.Atoi(parts[2])
			data, _ := json.Marshal(books[index])
			w.Write(data)
		}
	})

	http.ListenAndServe(":8080", nil)
}
