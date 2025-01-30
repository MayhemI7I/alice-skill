package main

import (
	"local/alice-skill/handlers"
	"local/alice-skill/internal/urlstorage"
	"net/http"
)

var Mux *http.ServeMux
var storage = urlstorage.NewURLStorage()

func main() {

	Mux = http.NewServeMux()
	Mux.HandleFunc("/", handlers.HandleURL(storage))

	run()

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	return http.ListenAndServe(":8086", Mux)
}

