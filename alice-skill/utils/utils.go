package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func GenerateShortURL(longurl string) string {
	hash := sha256.Sum256([]byte(longurl))
	shortURL := base64.URLEncoding.EncodeToString(hash[:])
	return shortURL[:8]

}
