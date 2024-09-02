package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type URL struct {
	OriginalURL  string `json:"original_url"`
	ShortenedURL string `json:"shortened_url"`
}

var urlStore = make(map[string]string)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/shorten", shortenURL).Methods("POST")
	router.HandleFunc("/{id}", redirectURL).Methods("GET")

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func shortenURL(w http.ResponseWriter, r *http.Request) {
	var url URL
	_ = json.NewDecoder(r.Body).Decode(&url)

	id := uuid.NewV4().String()[:8]
	shortenedURL := "http://localhost:8080/" + id

	urlStore[id] = url.OriginalURL

	url.ShortenedURL = shortenedURL
	json.NewEncoder(w).Encode(url)
}

func redirectURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	originalURL, ok := urlStore[id]
	if ok {
		http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
	} else {
		http.NotFound(w, r)
	}
}
