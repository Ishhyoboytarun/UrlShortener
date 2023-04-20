package repo

import (
	"net/http"

	error2 "UrlShortener/error"
	"UrlShortener/validator"
)

type UrlShortener interface {
	StoreUrl(longUrl string, shortUrl string) error
	GetShortUrl(longUrl string) (*validator.ShortUrlResponse, error)
}

type UrlShortenerRepo struct {
	longToShort map[string]string
	shortToLong map[string]string
}

func NewShortenUrlRepo(longToShort, shortToLong map[string]string) *UrlShortenerRepo {
	return &UrlShortenerRepo{
		longToShort: longToShort,
		shortToLong: shortToLong,
	}
}

type ShortUrlResponse struct {
	ShortUrl string `json:"short_url"`
}

func (s *UrlShortenerRepo) StoreUrl(longUrl string, shortUrl string) error {
	if _, ok := s.longToShort[longUrl]; ok {
		return nil // Already exists
	}

	if _, ok := s.shortToLong[shortUrl]; ok {
		return error2.ErrCollision
	}

	s.longToShort[longUrl] = shortUrl
	s.shortToLong[shortUrl] = longUrl
	return nil
}

func (s *UrlShortenerRepo) GetShortUrl(longUrl string) (*validator.ShortUrlResponse, error) {
	if shortUrl, ok := s.longToShort[longUrl]; ok {
		response := validator.NewShortUrlResponse()
		response.ShortUrl = shortUrl
		response.StatusCode = http.StatusOK
		return response, nil
	}

	return nil, error2.ErrNotFound
}
