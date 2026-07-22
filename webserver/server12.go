package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/chek", func(w http.ResponseWriter, r *http.Request) {

		strnum := r.URL.Query().Get("num")
		if strnum == "" {
			fmt.Fprint(w, "введите число")
			return
		}
		num, _ := strconv.Atoi(strnum)
		if num%2 == 0 {
			fmt.Fprint(w, "Четное")
		} else if num%2 != 0 {
			fmt.Fprint(w, "Нечетное")
		}
	})
	http.ListenAndServe(":8080", nil)

}
