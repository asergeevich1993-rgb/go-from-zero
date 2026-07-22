package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Nums struct {
	Value int
}

var data = map[string]int{"X": 1, "Y": 2}

func main() {

	http.HandleFunc("/data/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "PUT" {
			parts := strings.Split(r.URL.Path, "/")
			key := parts[2]

			body, _ := io.ReadAll(r.Body)
			defer r.Body.Close()

			var n Nums
			json.Unmarshal(body, &n)
			data[key] = n.Value
			fmt.Fprint(w, "Saved")

		}
		if r.Method == "GET" {
			jdata, _ := json.Marshal(data)
			w.Write(jdata)
		}
		if r.Method == "DELETE" {
			parts := strings.Split(r.URL.Path, "/")
			key := parts[2]

			delete(data, key)
			fmt.Fprint(w, "Deleted")
		}

	})
	http.ListenAndServe(":8080", nil)
}
