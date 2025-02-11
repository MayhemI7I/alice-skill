package main

import (
	"local/alice-skill/compression/zstd"
	"local/alice-skill/config"
	"local/alice-skill/handlers"
	"local/alice-skill/internal/urlstorage"
	"local/alice-skill/logger"
	"net/http"
)

func main() {

	cfg := config.InitConfig()

	logger.InitLogger(cfg.LogLevel)
	defer logger.CloseLogger()

	store := urlstorage.NewURLStorage()

	urlHandler := handlers.NewURLHandler(store)

	mux := http.NewServeMux()
	mux.Handle("/", handlers.WithLog(zstd.ZstdDecompress(zstd.ZstdCompress(http.HandlerFunc(urlHandler.HandURL)))))

	if err := run(cfg, mux); err != nil {
		logger.Log.Fatal(err)
	}
}

func run(cfg *config.Config, mux *http.ServeMux) error {
	logger.Log.Infof("Server started on %s:%s", cfg.ServerAdress, cfg.ServerPort)
	return http.ListenAndServe(cfg.ServerAdress+":"+cfg.ServerPort, mux)
}
