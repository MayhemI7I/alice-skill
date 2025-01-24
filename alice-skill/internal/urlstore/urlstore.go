package urlstore

import (
	"sync"
)

type URLStore struct {
	urls map[string]string
	mu sync.Mutex
}

func NewURLStore() *URLStore {
	return &URLStore{
		urls: make(map[string]string),
	mu: sync.Mutex{},
	}
}


 func (us *URLStore) SaveURL(shortUrl, longUrl string) {
	us.mu.Lock()
	defer us.mu.Unlock()
	us.urls[shortUrl] = longUrl
 }

 func (us *URLStore) GetURL(shortUrl string) string {
	us.mu.Lock()
	defer us.mu.Unlock()
	return us.urls[shortUrl]
 
 } 