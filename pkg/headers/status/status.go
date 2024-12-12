package status

import "net/http"

// StatusHTTP custom struct for show response status code in log middleware
type StatusHTTP struct {
	http.ResponseWriter
	StatusCode int
}

// WriteHeader method writes status code in http header
func (sr *StatusHTTP) WriteHeader(statusCode int) {
	sr.StatusCode = statusCode
	sr.ResponseWriter.WriteHeader(statusCode)
}

// NewStatusHTTP function init new StatusHTTP
func NewStatusHTTP(w http.ResponseWriter) *StatusHTTP {
	return &StatusHTTP{w, http.StatusOK}
}
