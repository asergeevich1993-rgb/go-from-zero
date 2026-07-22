package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Account struct {
	Username string
	Password string
}
type Pass struct {
	Status string
}

func main() {

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		body, _ := io.ReadAll(r.Body)

		defer r.Body.Close()

		var a Account
		json.Unmarshal(body, &a)

		if a.Password == "12345" {
			p := Pass{Status: "ok"}
			data, _ := json.Marshal(p)
			w.Write(data)
		} else {

			p := Pass{Status: "unauthorizated"}
			data, _ := json.Marshal(p)
			w.WriteHeader(401)
			w.Write(data)
		}

	})

	http.ListenAndServe(":8080", nil)
}
