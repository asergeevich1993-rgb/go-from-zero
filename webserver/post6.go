package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Echo struct {
	Text string
}

func main() {

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Ошибка чтения", err)
			return
		}

		var e Echo
		json.Unmarshal(body, &e)

		data, _ := json.Marshal(e)

		w.Write(data)

	})

	http.ListenAndServe(":8080", nil)

}
