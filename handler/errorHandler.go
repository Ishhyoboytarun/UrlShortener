package handler

import (
	"net/http"

	error2 "UrlShortener/error"
	"UrlShortener/validator"
)

func handleError(w http.ResponseWriter, err error) {

	if err == error2.ErrInvalidBody || err == error2.ErrUrl || err == error2.ErrInsecureUrl {
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	writeErrorResponse(w, http.StatusOK, err)
}

func writeErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	er := err.Error()
	body := validator.NewShortUrlResponse()
	body.Error = &er
	body.StatusCode = statusCode
	writeResponse(w, statusCode, body)
}
