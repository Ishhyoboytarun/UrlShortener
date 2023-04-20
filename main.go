package main

import (
	"net/http"

	"UrlShortener/handler"
	"UrlShortener/repo"
	"UrlShortener/service"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	handler := initializeDependencies()

	// POST /shorten
	r.HandleFunc("/getShortenUrl", handler.UrlShortenerHandler.ShortenUrl).Methods("POST")

	http.ListenAndServe(":8001", r)
}
