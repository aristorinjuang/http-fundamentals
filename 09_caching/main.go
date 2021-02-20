package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	css, err := ioutil.ReadFile("./assets/app.css")
	if err != nil {
		log.Fatal(err)
	}

	html, err := ioutil.ReadFile("./html/index.html")
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/app.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		w.Header().Set("Cache-Control", "public, max-age=604800, immutable")
		w.Header().Set("ETag", "v0")
		w.Write(css)
	})
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(html)
	})
	http.ListenAndServe(":80", r)
}
