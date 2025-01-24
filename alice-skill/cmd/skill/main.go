package main

import (
	"local/alice-skill/handlers"
	"local/alice-skill/internal/urlstore"
	"net/http"
)

var Mux *http.ServeMux
var storage = urlstore.NewURLStore()

func main() {

	Mux = http.NewServeMux()
	Mux.HandleFunc("/", mwPost(http.HandlerFunc(handlers.HandlerPost)))
	Mux.HandleFunc("/", mwGet(http.HandlerFunc(handlers.HandlerGet)))

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	return http.ListenAndServe(":8080", Mux)
}
