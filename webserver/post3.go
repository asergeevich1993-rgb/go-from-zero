package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DivRequest struct {
	A, B int
}
type DivResponse struct {
	Result int
}

func main() {

	http.HandleFunc("/divide", func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Ошибка чтения ", err)
			return
		}
		defer r.Body.Close()

		var d DivRequest
		json.Unmarshal(body, &d)
		if d.B == 0 {
			w.WriteHeader(400)
			fmt.Fprint(w, "деление на 0")
			return
		}
		resp := DivResponse{Result: d.A / d.B}

		data, _ := json.Marshal(resp)

		w.Write(data)
	})

	http.ListenAndServe(":8080", nil)
}
