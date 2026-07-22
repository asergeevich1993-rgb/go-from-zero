package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Nums struct {
	Items string
}

var items = []string{"один", "два", "три"}

func main() {

	http.HandleFunc("/items/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "PUT" {
			parts := strings.Split(r.URL.Path, "/")
			index, _ := strconv.Atoi(parts[2])

			body, _ := io.ReadAll(r.Body)
			defer r.Body.Close()

			var n Nums
			json.Unmarshal(body, &n)
			items[index] = n.Items
			fmt.Fprint(w, "Обновлено")

		}
		if r.Method == "GET" {
			data, _ := json.Marshal(items)

			w.Write(data)
		}
		if r.Method == "DELETE" {
			parts := strings.Split(r.URL.Path, "/")
			index, _ := strconv.Atoi(parts[2])
			items = append(items[:index], items[index+1:]...)
			fmt.Fprint(w, "удалено")
		}

	})
	http.ListenAndServe(":8080", nil)
}
