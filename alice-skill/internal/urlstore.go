package internal

import (
	"sync"
)

 var (
	URLStore = make(map[string]string)
	mu sync.Mutex
 )

 func SaveURL(shortUrl, longUrl string) {
	mu.Lock()
	defer mu.Unlock()
	URLStore[shortUrl] = longUrl
 }

 func GetURL(shortUrl string) string {
	mu.Lock()
	defer mu.Unlock()
	return URLStore[shortUrl]
 
 }