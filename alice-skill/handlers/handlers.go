package handlers

import (
	"fmt"
	"net/http"
	"local/alice-skill/utils"
	"local/alice-skill/internal/urlstore"

	
)

func parsURL(w http.ResponseWriter, r *http.Request)string {
	shortURL := r.URL.Path[1:]
	longURL := urlstore.GetURL(shortURL)
	if longURL == "" {
		http.Error( w,"Long URL is not exist",400)
		return ""
	}
	return longURL
}
		



func HandlerGet(s urlstore.URLStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			longURL := parsURL(w,r)
			if longURL == "" {
				http.Error(w, "Long URL is not exist", 400)
				return
			}
			http.Redirect(w, r, longURL, 307)
	
		case "POST":
			var longURLPost string
			_, err := fmt.Fscan(r.Body, longURLPost)
			if err != nil || longURLPost == "" || longURLPost == " " {
				http.Error(w, "Invalid URL", 400)
				return
			}
			shortURL := utils.GenerateShortURL(longURLPost)
			s.SaveURL(shortURL, longURLPost)
		default:
			http.Error(w, "Method not allowed", 400)
			return
			}	



	}

}

func HandlerPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", 400)
		return
	}
	var longUrl string 
	 _, err := fmt.Fscan(r.Body, longUrl)
	 if err != nil || longUrl == "" || longUrl == " " {
		http.Error(w, "Invalid URL", 400)
		 return
		}
	shortURL := utils.GenerateShortURL(longUrl)
	urlstore.SaveURL(shortURL, longUrl)
	w.WriteHeader(201)
	w.Write([]byte(shortURL))		 
}