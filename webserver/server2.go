package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Главная страница") })
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "О нас") })
	http.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Контакты") })

	http.ListenAndServe(":8080", nil)
}
