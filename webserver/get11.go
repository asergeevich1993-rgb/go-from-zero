package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Divisor struct {
	A, B int
	Div  float64
}

func main() {

	http.HandleFunc("/div", func(w http.ResponseWriter, r *http.Request) {

		astr := r.URL.Query().Get("a")
		bstr := r.URL.Query().Get("b")

		a, _ := strconv.Atoi(astr)
		b, _ := strconv.Atoi(bstr)
		if b == 0 {
			fmt.Fprint(w, "деление на ноль")
			return
		}

		nums := Divisor{A: a, B: b, Div: float64(a) / float64(b)}

		data, _ := json.Marshal(nums)

		w.Write(data)

	})

	http.ListenAndServe(":8080", nil)
}
