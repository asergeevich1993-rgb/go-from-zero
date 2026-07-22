package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Multi struct {
	X       int
	Y       int
	Product int
}

func main() {

	http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
		strx := r.URL.Query().Get("x")
		stry := r.URL.Query().Get("y")

		x, _ := strconv.Atoi(strx)
		y, _ := strconv.Atoi(stry)

		nums := Multi{X: x, Y: y, Product: x * y}
		data, _ := json.Marshal(nums)
		w.Write(data)

	})
	http.ListenAndServe(":8080", nil)

}
