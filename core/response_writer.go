package core

import "net/http"

// ResponseWriter represents an HTTP response.
type ResponseWriter interface {
	http.ResponseWriter
	// Status returns the status code of the response or 0 if the response has not been written.
	Status() int
}
