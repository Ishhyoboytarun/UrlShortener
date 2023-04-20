package handler

import (
	"encoding/json"
	"net/http"

	"UrlShortener/service"
	"UrlShortener/validator"
)

type UrlShortenerHandler struct {
	service service.UrlShortener
}

func NewShortenUrlHandler(service service.UrlShortener) *UrlShortenerHandler {
	return &UrlShortenerHandler{
		service: service,
	}
}

func (s UrlShortenerHandler) ShortenUrl(w http.ResponseWriter, request *http.Request) {

	shortenUrlRequest := &validator.ShortenUrlRequest{}
	if err := json.NewDecoder(request.Body).Decode(&shortenUrlRequest); err != nil {
		handleError(w, err)
		return
	}

	err := shortenUrlRequest.Validate()
	if err != nil {
		handleError(w, err)
		return
	}

	shortenUrlResponse, err := s.service.ShortenUrl(shortenUrlRequest)
	if err != nil {
		handleError(w, err)
		return
	}

	writeResponse(w, http.StatusOK, shortenUrlResponse)
}
