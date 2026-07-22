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
	Name string
}

var users = []string{"Мария", "Артур", "Иван"}

func main() {

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "PUT" {
			parts := strings.Split(r.URL.Path, "/")
			index, err := strconv.Atoi(parts[2])
			if err != nil {
				fmt.Fprintf(w, "Ошибка конвертации %v", err)
			}

			body, _ := io.ReadAll(r.Body)
			defer r.Body.Close()

			var u UserRequest
			json.Unmarshal(body, &u)
			users[index] = u.Name
			fmt.Fprint(w, "Добавлено")
		} else if r.Method == "GET" {

			data, _ := json.Marshal(users)

			w.Write(data)

		}

	})
	http.ListenAndServe(":8080", nil)

}
