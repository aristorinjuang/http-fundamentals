package main

import (
	"io"
	"net/http"
)

func respond(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<!DOCTYPE html><head><title>Hello World!</title></head><body><h1>Hello World!</h1></body></html>")
}

func main() {
	http.HandleFunc("/", respond)
	http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
}
