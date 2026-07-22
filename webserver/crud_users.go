package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type UserRequest struct {
	Users string
}

var users []string

func main() {

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Ошибка чтения %v", err)
				return
			}
			defer r.Body.Close()
			var newuser UserRequest

			json.Unmarshal(body, &newuser)

			users = append(users, newuser.Users)

			data, _ := json.Marshal(users)

			w.Write(data)

		}
		if r.Method == "GET" {
			data, _ := json.Marshal(users)
			w.Write(data)
		}
	})
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "PUT" {
			parts := strings.Split(r.URL.Path, "/")
			index, err := strconv.Atoi(parts[2])
			if err != nil {
				fmt.Fprint(w, "Ошибка конвертации")
				return
			}
			body, _ := io.ReadAll(r.Body)
			defer r.Body.Close()

			var newuser UserRequest
			json.Unmarshal(body, &newuser)
			users[index] = newuser.Users
			fmt.Fprint(w, "Обновлено")
		}
		if r.Method == "DELETE" {
			parts := strings.Split(r.URL.Path, "/")
			index, _ := strconv.Atoi(parts[2])

			users = append(users[:index], users[index+1:]...)
			fmt.Fprint(w, "удалено")
		}
		if r.Method == "GET" {
			parts := strings.Split(r.URL.Path, "/")
			index, _ := strconv.Atoi(parts[2])
			data, _ := json.Marshal(users[index])

			w.Write(data)
		}

	})
	http.ListenAndServe(":8080", nil)
}
