package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Nums struct {
	A, B int
}

type Result struct {
	Sum int
}

func main() {

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {

		body, _ := io.ReadAll(r.Body)

		defer r.Body.Close()

		var n Nums
		json.Unmarshal(body, &n)

		result := Result{Sum: n.A + n.B}

		data, _ := json.Marshal(result)

		w.Write(data)

	})
	http.ListenAndServe(":8080", nil)

}
