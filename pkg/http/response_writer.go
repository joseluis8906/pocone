package http

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	code int
	size int
}

// newResponseWriter returns a new ResponseWriter
func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
	}
}

// WriteHeader overrides the WriteHeader method to store the status code
func (r *responseWriter) WriteHeader(code int) {
	if r.Code() == 0 {
		r.code = code
	}
	r.ResponseWriter.WriteHeader(code)
}

// Write overrides the Write method to store the response size
func (r *responseWriter) Write(body []byte) (int, error) {
	if r.Code() == 0 {
		r.WriteHeader(http.StatusOK)
	}
	var err error
	r.size, err = r.ResponseWriter.Write(body)
	return r.size, err
}

// Flush overrides the Flush method if the ResponseWriter implements it
func (r *responseWriter) Flush() {
	if fl, ok := r.ResponseWriter.(http.Flusher); ok {
		if r.Code() == 0 {
			r.WriteHeader(http.StatusOK)
		}
		fl.Flush()
	}
}

// Hijack overrides the Hijack method if the ResponseWriter implements it
func (r *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hj, ok := r.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, fmt.Errorf("the hijacker interface is not supported")
	}
	return hj.Hijack()
}

// Code returns the stored response status code
func (r *responseWriter) Code() int {
	return r.code
}

// Size returns the stored response size
func (r *responseWriter) Size() int {
	return r.size
}
