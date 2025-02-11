package urlstorage

import (
	"errors"
	"local/alice-skill/logger"
	"log"
	"sync"
)



type URLStorage struct {
	urls map[string]string 
	mu   sync.Mutex
}

func NewURLStorage() *URLStorage {
	return &URLStorage{
		urls: make(map[string]string),
		mu:   sync.Mutex{},
	}
}

func (us *URLStorage) SaveURL(shortURL, longURL string) error {
	us.mu.Lock()
	defer us.mu.Unlock()
	if shortURL == "" || longURL == "" {
		log.Printf("Invalid argument: %s, %s", shortURL, longURL)
		return errors.New("invalid argument")
	}
	if _, exists := us.urls[shortURL]; exists {
		log.Printf("URL already exists: %s", shortURL)
		return errors.New("URL already exists")
	}
	us.urls[shortURL] = longURL
	logger.Log.Info("Saved: %s -> %s", shortURL, longURL)
	return nil
}


func (us *URLStorage) GetURL(shortUrl string) (string, error) {
	us.mu.Lock()
	defer us.mu.Unlock()
	if shortUrl == "" {
		log.Printf("Invalid argument: %s", shortUrl)
		return "", errors.New("invalid short URL argument")
  
	}
	value, ok := us.urls[shortUrl]
	if !ok {
		return "", errors.New("URL not found in storage")
	}
	logger.Log.Info("Retrived: %s -> %s", shortUrl, value)
	return value, nil

}
