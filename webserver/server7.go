package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Time string
}

func main() {

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now().UTC().Format(time.RFC3339)
		t := Response{Time: now}

		data, _ := json.Marshal(t)

		w.Write(data)
	})
	http.ListenAndServe(":8080", nil)

}
