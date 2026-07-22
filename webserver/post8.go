package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

type isAdult struct {
	Name  string
	Age   int
	Adult bool
}

func main() {

	http.HandleFunc("/isadult", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)

		defer r.Body.Close()

		var p Person

		json.Unmarshal(body, &p)

		if p.Age >= 18 {
			a := isAdult{Name: p.Name, Age: p.Age, Adult: true}
			data, _ := json.Marshal(a)
			w.Write(data)
		} else {
			a := isAdult{Name: p.Name, Age: p.Age, Adult: false}
			data, _ := json.Marshal(a)
			w.Write(data)
		}

	})
	http.ListenAndServe(":8080", nil)

}
