// Package clienthelper provides utility functions for creating API responses.

package clienthelper

import (
	"encoding/json"
	"net/http"
)

// APIResponse is a struct representing the response to be sent back to the client.
type APIResponse struct {
	code     int
	headers  http.Header
	writer   http.ResponseWriter
	response Response
}

// NewAPIResponse creates a new instance of APIResponse with an empty header.
func NewAPIResponse(writer http.ResponseWriter) *APIResponse {
	return &APIResponse{
		writer:  writer,
		headers: make(http.Header),
	}
}

// SetStatusCode sets the status code of the APIResponse and returns a pointer to the APIResponse.
func (r *APIResponse) SetStatusCode(code int) *APIResponse {
	r.code = code
	return r
}

// SetHeader sets the HTTP header of the APIResponse and returns a pointer to the APIResponse.
func (r *APIResponse) SetHeader(key, value string) *APIResponse {
	r.headers.Set(key, value)
	return r
}

// GetStatusCode gets the status code of the APIResponse.
func (r *APIResponse) GetStatusCode() int {
	return r.code
}

// GetHeader gets the HTTP header value of the APIResponse.
func (r *APIResponse) GetHeader(key string) string {
	return r.headers.Get(key)
}

// GetHeaders gets all the HTTP headers of the APIResponse.
func (r *APIResponse) GetHeaders() http.Header {
	return r.headers
}

// GetResponse gets the response body of the APIResponse.
func (r *APIResponse) GetResponse() *Response {
	return &r.response
}

func (r *APIResponse) SetResponse(resp Response) {
	r.response = resp
}

// Send writes the HTTP header and response body to the response writer.
func (r *APIResponse) Send() error {
	if r.code == 0 {
		r.code = http.StatusOK
	}

	// Write the headers to the response writer
	for key, values := range r.headers {
		for _, value := range values {
			r.writer.Header().Add(key, value)
		}
	}

	// Write the status code to the response writer
	r.writer.WriteHeader(r.code)

	// Write the response body to the response writer
	err := json.NewEncoder(r.writer).Encode(r.response)

	return err
}
