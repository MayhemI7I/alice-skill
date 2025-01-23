package main

import (
	"local/alice-skill/handlers"
	"net/http"
)

var Mux *http.ServeMux

func main() {
	Mux = http.NewServeMux()
	Mux.HandleFunc("/", http.HandlerFunc(handlers.HandlerPost))
	Mux.HandleFunc("/get", http.HandlerFunc(handlers.HandlerGet))

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	return http.ListenAndServe(":8080", Mux)
}
