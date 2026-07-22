package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type DataRequest struct {
	Key   string
	Value int
}

var data = map[string]int{}

func main() {

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			body, _ := io.ReadAll(r.Body)

			defer r.Body.Close()

			var d DataRequest
			json.Unmarshal(body, &d)

			data[d.Key] = d.Value
			jdata, _ := json.Marshal(data)
			w.Write(jdata)
		}
		if r.Method == "GET" {
			jdata, _ := json.Marshal(data)
			w.Write(jdata)
		}

	})
	http.HandleFunc("/data/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			parts := strings.Split(r.URL.Path, "/")
			key := parts[2]
			body, _ := io.ReadAll(r.Body)
			defer r.Body.Close()

			var d DataRequest
			json.Unmarshal(body, &d)

			data[key] = d.Value
			fmt.Fprint(w, "Обновлено")

		}
		if r.Method == "DELETE" {
			parts := strings.Split(r.URL.Path, "/")
			key := parts[2]

			delete(data, key)
			fmt.Fprint(w, "удалено")
		}
		if r.Method == "GET" {
			parts := strings.Split(r.URL.Path, "/")
			key := parts[2]
			jdata, _ := json.Marshal(data[key])
			w.Write(jdata)
		}
	})
	http.ListenAndServe(":8080", nil)

}
