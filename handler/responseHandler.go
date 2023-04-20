package handler

import (
	"encoding/json"
	"net/http"
)

func writeResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	writeResponseWithOptions(w, statusCode, body, true)
}

func writeResponseWithOptions(w http.ResponseWriter, statusCode int, body interface{}, encodeHTML bool) {
	w.Header().Set(headerContentType, contentTypeJSON)
	w.Header().Set(headerAccessControlAllowOrigin, headerValueAll)
	w.Header().Set(headerAccessControlAllowHeaders, headerValueAll)
	w.Header().Set(headerAccessControlAllowMethods, headerValueAllowedMethod)
	w.WriteHeader(statusCode)

	if body == nil || statusCode == http.StatusNoContent {
		return
	}

	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(encodeHTML)

	_ = encoder.Encode(body)
	return
}

const (
	headerContentType               = "Content-Type"
	headerAccessControlAllowOrigin  = "Access-Control-Allow-Origin"
	headerAccessControlAllowHeaders = "Access-Control-Allow-Headers"
	headerAccessControlAllowMethods = "Access-Control-Allow-Methods"
	headerValueAll                  = "*"
	headerValueAllowedMethod        = "POST"
	contentTypeJSON                 = "application/json"
)
