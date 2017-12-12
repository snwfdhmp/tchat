package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	Messages = make([]string, 0)
)

func write(w http.ResponseWriter, v interface{}) {
	b, _ := json.Marshal(v)
	w.Write(b)
}

func read(r *http.Request) (string, error) {
	var s string
	return s, json.NewDecoder(r.Body).Decode(&s)
}

func headers(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		h.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		write(w, Messages)

		Messages = make([]string, 0)
	}).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg, err := read(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		Messages = append(Messages, msg)
		log.Println("Received '", msg, "' at", time.Now())
	}).Methods("POST")

	http.ListenAndServe(":9876", headers(r))
}
