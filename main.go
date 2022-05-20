package main

import (
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/pipeline", pipeline)
	http.ListenAndServe(":8080", nil)
	// http.ListenAndServeTLS( ":8443", "./cert.pem", "./key.pem", nil )
}

var array []string

func pipeline(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	add, okForm := r.Form["add"]
	if okForm && len(add) == 1 {
		array = append(array, string(add[0]))
	}

	id, ok := r.URL.Query()["id"]
	if ok && len(id) == 1 {
		id, err := strconv.Atoi(id[0])
		if err != nil {
			return
		}
		for i, info := range array {
			if i == id {
				w.Write([]byte(info))
				return
			}
		}
	}
}
