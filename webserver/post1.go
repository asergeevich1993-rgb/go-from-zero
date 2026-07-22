package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Ошибка чтения", err)
		}
		defer r.Body.Close()

		var p Person
		json.Unmarshal(body, &p)

		fmt.Fprintf(w, "Привет %s, тебе %d лет", p.Name, p.Age)

	})
	http.ListenAndServe(":8080", nil)

}
