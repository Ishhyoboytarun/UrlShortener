package service

import (
	"strconv"
	"time"

	error2 "UrlShortener/error"
	"UrlShortener/repo"
	"UrlShortener/validator"
)

type UrlShortener interface {
	ShortenUrl(shortenUrlRequest *validator.ShortenUrlRequest) (*validator.ShortUrlResponse, error)
}

type UrlShortenerService struct {
	repo repo.UrlShortener
}

func NewShortenUrlService(repo repo.UrlShortener) *UrlShortenerService {
	return &UrlShortenerService{
		repo: repo,
	}
}

func (s *UrlShortenerService) ShortenUrl(shortenUrlRequest *validator.ShortenUrlRequest) (*validator.ShortUrlResponse,
	error) {

	// Generate short URL using timestamp-based encoding
	var shortUrl string
	for i := 0; i < 100; i++ { // Try 100 times to avoid collisions
		shortUrl = "https://bit.ly/" + strconv.FormatInt(time.Now().UnixNano(), 32)[:10]
		if err := s.repo.StoreUrl(shortenUrlRequest.LongUrl, shortUrl); err == nil {
			break // Success!
		} else if err == error2.ErrCollision {
			continue // Try again
		} else {
			return nil, err
		}
	}
	return s.repo.GetShortUrl(shortenUrlRequest.LongUrl)
}
