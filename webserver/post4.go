package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Person struct {
	Name string
}
type Greeting struct {
	Greet string
}

func main() {

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("ошибка чтения", err)
			return
		}
		defer r.Body.Close()

		var p Person
		json.Unmarshal(body, &p)

		if p.Name == "" {
			p.Name = "гость"

		}
		gr := Greeting{Greet: "Привет " + p.Name + "!"}

		data, _ := json.Marshal(gr)

		w.Write(data)

	})
	http.ListenAndServe(":8080", nil)

}
