package http

import (
	"net/http"

	"github.com/go-chi/render"
)

// ErrResponse represents an error response
type ErrResponse struct {
	Err            error  `json:"-"`               // low-level runtime error
	HTTPStatusCode int    `json:"-"`               // http response status code
	AppCode        int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText      string `json:"error,omitempty"` // application-level error message
}

// Render sets the HTTP status code and writes the error response
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrInternal creates an internal server error response
func ErrInternal(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		ErrorText:      err.Error(),
	}
}

// ErrBadRequest creates a bad request error response
func ErrBadRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		ErrorText:      err.Error(),
	}
}
