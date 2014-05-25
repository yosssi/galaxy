package core

import "net/http"

// response represents an HTTP response.
type response struct {
	rw     http.ResponseWriter
	status int
}

// Header invokes http.ResponseWriter.Header.
func (res *response) Header() http.Header {
	return res.rw.Header()
}

// Write invokes http.ResponseWriter.Write.
func (res *response) Write(b []byte) (int, error) {
	if res.status == 0 {
		res.status = http.StatusOK
	}
	return res.rw.Write(b)
}

// WriteHeader invokes http.ResponseWriter.WriteHeader.
func (res *response) WriteHeader(i int) {
	res.status = i
	res.rw.WriteHeader(i)
}

// Status returns the response's status.
func (res *response) Status() int {
	return res.status
}

// newResponse generates a response and returns it.
func newResponse(rw http.ResponseWriter) *response {
	return &response{
		rw: rw,
	}
}
