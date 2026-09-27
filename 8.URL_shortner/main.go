package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalUrl  string    `json:"original_url"`
	ShortUrl     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var UrlDB = make(map[string]URL)

func generateShortUrl(OriginalUrl string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalUrl)) // convert original url string to a byte slice
	//fmt.Println("hasher:", hasher)
	data := hasher.Sum(nil)
	//fmt.Println("hasher data: ", data)
	hash := hex.EncodeToString(data)
	//fmt.Println("Encoding to String: ", hash)
	//fmt.Println("final string: ", hash[:8])
	return hash[:8]
}
func createURL(originalUrl string) string {
	shortURL := generateShortUrl(originalUrl)
	id := shortURL
	UrlDB[id] = URL{
		ID:           id,
		OriginalUrl:  originalUrl,
		ShortUrl:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}

func getUrl(id string) (URL, error) {
	url, ok := UrlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil

}
func RootpageUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Word !!")

}
func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortURL_ := createURL(data.URL)
	// fmt.Fprintf(w, shortURL)
	response := struct {
		ShortURL string `json:"short_url"`
	}{ShortURL: shortURL_}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := getUrl(id)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.OriginalUrl, http.StatusFound)
}

func main() {
	// fmt.Println("Starting URL shortener...")
	// OriginalUrl := "https://docs-cybersec.thalesgroup.com/bundle/on-premises-knowledgebase-reference-guide/page/abnormally_long_url.htm"
	// generateShortURL(OriginalUrl)

	// Register the handler function to handle all requests to the root URL ("/")
	http.HandleFunc("/", RootpageUrl)
	http.HandleFunc("/shorten", ShortURLHandler)
	http.HandleFunc("/redirect/", redirectURLHandler)

	// Start the HTTP server on port 8080
	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error on starting server:", err)
	}
}
