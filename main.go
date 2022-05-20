package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/pipelines", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getPipelines(w, r)

		case http.MethodPost:
			addPipelines(w, r)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Metodo no permitido")
			return
		}
	})
	http.ListenAndServe(":8080", nil)
	// http.ListenAndServeTLS( ":8443", "./cert.pem", "./key.pem", nil )
}

type Pipes struct {
	Name   string
	Status string
}

var pipelines []*Pipes = []*Pipes{}

func getPipelines(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pipelines)
}

func addPipelines(w http.ResponseWriter, r *http.Request) {
	pipes := &Pipes{}
	err := json.NewDecoder(r.Body).Decode(pipes)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}
	pipelines = append(pipelines, pipes)
}
