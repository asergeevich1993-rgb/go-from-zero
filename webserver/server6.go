package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Ping struct {
	Status string
}

func main() {

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		s := Ping{Status: "ok"}
		data, err := json.Marshal(s)
		if err != nil {
			fmt.Println("Ошибка сериализации")
			return
		}
		w.Write(data)

	})
	http.ListenAndServe(":8080", nil)
}
