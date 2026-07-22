package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Nums struct {
	A, B int
}
type Maximum struct {
	Max int
	Msg string
}

func main() {

	http.HandleFunc("/max", func(w http.ResponseWriter, r *http.Request) {

		body, _ := io.ReadAll(r.Body)

		defer r.Body.Close()

		var n Nums
		json.Unmarshal(body, &n)

		if n.A == n.B {
			m := Maximum{Msg: "Равны"}
			data, _ := json.Marshal(m)
			w.Write(data)
			return

		}
		max1 := n.A
		if n.B > n.A {
			max1 = n.B
		}
		m := Maximum{Max: max1}
		data, _ := json.Marshal(m)

		w.Write(data)

	})
	http.ListenAndServe(":8080", nil)
}
