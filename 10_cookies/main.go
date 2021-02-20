package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Set-Cookie", "foo=bar; Max-Age=86400")

		foo, err := r.Cookie("foo")
		if err != nil {
			log.Fatal(err)
		}

		w.Write([]byte("foo = " + foo.Value))
	})
	http.ListenAndServe(":80", r)
}
