package httperrors

import (
	"net/http"
)

// HTTPError is the custom error type that wraps common http errors
type HTTPError struct {
	Request *http.Request `json:"-"`
	Error   error         `json:"-"`
	Message string        `json:"error"`
	Code    int           `json:"status"`
}

// Handler is a function that handles an http request and returns the custom error type HttpError
// when an error occurs
type Handler func(w http.ResponseWriter, r *http.Request) *HTTPError
