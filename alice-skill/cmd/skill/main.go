package main

import (
	"local/alice-skill/handlers"
	"local/alice-skill/internal/urlstore"
	"net/http"
)

var Mux *http.ServeMux
var storage = urlstorage.NewURLStore()

func main() {

	Mux = http.NewServeMux()
	Mux.HandleFunc("/", handlers.HandleURL(storage))

	run()

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	return http.ListenAndServe(":8082", Mux)
}
