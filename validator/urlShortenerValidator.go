package validator

import (
	"net/http"
	"net/url"
	"regexp"
	"strings"

	err "UrlShortener/error"
)

type ShortenUrlRequest struct {
	LongUrl string `json:"long_url"`
}

type ShortUrlResponse struct {
	StatusCode int     `json:"status_code"`
	ShortUrl   string  `json:"short_url"`
	Error      *string `json:"error"`
}

func NewShortUrlResponse() *ShortUrlResponse {
	return &ShortUrlResponse{
		StatusCode: http.StatusBadRequest,
		ShortUrl:   "",
		Error:      nil,
	}
}

var urlRegex = regexp.MustCompile(`^(https?|ftp)://[^\s/$.?#].[^\s]*$`)

func (Url *ShortenUrlRequest) Validate() error {

	if Url.LongUrl == "" {
		return err.ErrEmptyUrl
	}

	if !strings.Contains(Url.LongUrl, "https") {
		return err.ErrInsecureUrl
	}

	// Check if the URL is in a valid format
	_, er := url.ParseRequestURI(Url.LongUrl)
	if er != nil {
		return err.ErrUrl
	}

	// Check if the URL matches the regular expression
	if !urlRegex.MatchString(Url.LongUrl) {
		return err.ErrUrl
	}

	return nil
}
