package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	Messages []string
)

func write(w http.ResponseWriter, v interface{}) {
	b, _ := json.Marshal(v)
	w.Write(b)
}

func read(r *http.Request) (string, error) {
	var s string
	return s, json.NewDecoder(r.Body).Decode(&s)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		write(w, Messages)
	}).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg, err := read(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		Messages = append(Messages, msg)
	}).Methods("POST")

	http.ListenAndServe(":9876", r)
}
