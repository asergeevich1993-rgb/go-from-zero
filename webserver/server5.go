package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		p := Person{Name: "Артур", Age: 33}

		data, err := json.Marshal(p)
		if err != nil {
			fmt.Println("ошибка сериализации", err)
			return
		}
		w.Write(data)
	})
	http.ListenAndServe(":8080", nil)
}
