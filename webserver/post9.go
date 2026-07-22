package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Parametrs struct {
	Width  int
	Height int
}
type Results struct {
	Area      int
	Perimeter int
}

func main() {

	http.HandleFunc("/rect", func(w http.ResponseWriter, r *http.Request) {

		body, _ := io.ReadAll(r.Body)

		defer r.Body.Close()

		var p Parametrs

		json.Unmarshal(body, &p)

		if p.Width == 0 || p.Height == 0 {
			w.WriteHeader(400)
			fmt.Fprint(w, "неверные размеры")
			return
		}
		req := Results{Area: p.Width * p.Height, Perimeter: (p.Width + p.Height) * 2}

		data, _ := json.Marshal(req)

		w.Write(data)

	})
	http.ListenAndServe(":8080", nil)

}
