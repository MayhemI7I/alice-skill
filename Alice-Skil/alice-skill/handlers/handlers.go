package handlers

import (
	"io"
	"local/alice-skill/internal/urlstorage"
	"local/alice-skill/utils"
	"net/http"
)


func parseURL(storage *urlstorage.URLStorage, w http.ResponseWriter, r *http.Request) string {
	shortURL := r.URL.Path[1:] // Извлекаем короткий URL из пути запроса
	longURL, err := storage.GetURL(shortURL)
	if err != nil || longURL == "" || longURL == " " {
		http.Error(w, "Long URL does not exist", http.StatusBadRequest) // Возвращаем ошибку, если длинный URL не найден
		return ""
	}

	return longURL
}

func handleGet(w http.ResponseWriter, r *http.Request, storage *urlstorage.URLStorage) {
	longURL := parseURL(storage, w, r) // Получаем длинный URL
	if longURL == "" {
		// Если длинного URL нет, отправляем ошибку и не продолжаем выполнение
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	// Отправляем редирект на длинный URL
	w.Header().Set("Location", longURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func handlePost(w http.ResponseWriter, r *http.Request, storage *urlstorage.URLStorage) {
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	longURLPost := string(body)
	if longURLPost == "" || longURLPost == " " {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	shortURL, err := utils.GenerateShortURL(longURLPost)
	if err != nil {
		http.Error(w, "Error generating short URL", http.StatusInternalServerError)
		return

	}
	storage.SaveURL(shortURL, longURLPost)
	w.Header().Set("Content-type", "text/plain; charset=utf-8")
	w.Write([]byte(shortURL))         // Отправляем сокращённый URL
	w.WriteHeader(http.StatusCreated) // Ответ с кодом 201
}

func HandleURL(storage *urlstorage.URLStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGet(w, r, storage)
		case http.MethodPost:
			handlePost(w, r, storage)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
