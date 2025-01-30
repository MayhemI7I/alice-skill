package urlstorage

import (
	"errors"
	"sync"
	"log"
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

func (us *URLStorage) SaveURL(shortURL, longURL string)  {
	us.mu.Lock()
	defer us.mu.Unlock()
	if shortURL == "" || longURL == "" {
		log.Printf("Invalid argument: %s, %s", shortURL, longURL)
		return 
	}
	if _, exists := us.urls[shortURL]; exists {
		log.Printf("URL already exists: %s", shortURL)
		return
	}
	us.urls[shortURL] = longURL
	log.Printf("Saved: %s -> %s", shortURL, longURL)
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
	log.Printf("Retrived: %s -> %s", shortUrl, value)
	return value, nil

}
