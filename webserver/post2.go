package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	Value int
}
type Response struct {
	Result int
}

func main() {

	http.HandleFunc("/double", func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Ошибка чтения:", err)
			return
		}
		defer r.Body.Close()

		var req Request
		json.Unmarshal(body, &req)

		res := Response{Result: req.Value * 2}

		data, _ := json.Marshal(res)

		w.Write(data)
	})
	http.ListenAndServe(":8080", nil)

}
