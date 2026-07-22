package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Add struct {
	A   int
	B   int
	Sum int
}

func main() {

	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {

		astr := r.URL.Query().Get("a")
		bstr := r.URL.Query().Get("b")
		a, _ := strconv.Atoi(astr)
		b, _ := strconv.Atoi(bstr)
		nums := Add{A: a, B: b, Sum: a + b}
		data, _ := json.Marshal(nums)
		w.Write(data)
	})
	http.ListenAndServe(":8080", nil)

}
