package main

import (
	"net/http"

	"UrlShortener/handler"
	"UrlShortener/repo"
	"UrlShortener/service"
	"github.com/gorilla/mux"
)

type Dependencies struct {
	UrlShortenerHandler *handler.UrlShortenerHandler
}

func main() {
	r := mux.NewRouter()

	handler := initializeDependencies()

	// POST /shorten
	r.HandleFunc("/getShortenUrl", handler.UrlShortenerHandler.ShortenUrl).Methods("POST")

	http.ListenAndServe(":8001", r)
}

func initializeDependencies() *Dependencies {
	longToShort := make(map[string]string)
	shortToLong := make(map[string]string)
	urlShortenerRepo := repo.NewShortenUrlRepo(longToShort, shortToLong)
	urlShortenerService := service.NewShortenUrlService(urlShortenerRepo)
	urlShortenerHandler := handler.NewShortenUrlHandler(urlShortenerService)
	return &Dependencies{
		UrlShortenerHandler: urlShortenerHandler,
	}
}
