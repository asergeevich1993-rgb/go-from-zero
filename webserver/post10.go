package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Word struct {
	Text string
}

type UpWord struct {
	Upper string
}

func main() {

	http.HandleFunc("/upper", func(w http.ResponseWriter, r *http.Request) {

		body, _ := io.ReadAll(r.Body)

		defer r.Body.Close()

		var ww Word
		json.Unmarshal(body, &ww)

		u := UpWord{Upper: strings.ToUpper(ww.Text)}

		data, _ := json.Marshal(u)

		w.Write(data)

	})

	http.ListenAndServe(":8080", nil)

}
