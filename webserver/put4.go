package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type UserScore struct {
	Score int
}

var scores = map[string]int{"Артур": 100, "Мария": 200}

func main() {

	http.HandleFunc("/scores/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "PUT" {
			parts := strings.Split(r.URL.Path, "/")
			name := parts[2]

			body, _ := io.ReadAll(r.Body)
			defer r.Body.Close()

			var req UserScore
			json.Unmarshal(body, &req)

			scores[name] = req.Score
			fmt.Fprint(w, "Изменено")
			fmt.Println(name, req.Score)
		}

		if r.Method == "GET" {

			data, _ := json.Marshal(scores)
			w.Write(data)
		}

	})
	http.ListenAndServe(":8080", nil)
}
