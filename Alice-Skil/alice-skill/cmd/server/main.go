package main

import (
	"local/alice-skill/config"
	"local/alice-skill/handlers"
	"local/alice-skill/internal/urlstorage"
	"log"
	"net/http"
)

var Mux *http.ServeMux
var storage = urlstorage.NewURLStorage()

func main() {
	cfg := config.InitConfig()

	Mux = http.NewServeMux()
	Mux.HandleFunc("/", handlers.HandleURL(storage))

	if err := run(cfg, Mux); err != nil {
		log.Fatal(err)
	}
	log.Printf("Server runed on %s:%s", cfg.ServerAdress, cfg.ServerPort)
}

func run(cfg *config.Config, mux *http.ServeMux) error {
	return http.ListenAndServe(cfg.ServerAdress+":"+cfg.ServerPort, mux)
}
