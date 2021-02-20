package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<!DOCTYPE html><head><title>Hello World!</title></head><body><h1>Hello World!</h1></body></html>"))
	})
	http.ListenAndServe(":80", r)
}
