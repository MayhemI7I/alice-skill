package handlers

import (
	"fmt"
	"net/http"
	"local/alice-skill/utils"

	
)


func HandlerGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
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

	w.Write([]byte(shortURL))	
	

	 
}