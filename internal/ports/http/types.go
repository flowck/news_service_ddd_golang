package http

import (
	"net/http"

	"github.com/flowck/news_service_ddd_golang/internal/ports/http/static"
)

type GenericResponse static.GenericResponse

type ErrResponse struct {
	Err     error  `json:"-"`
	Message string `json:"message,omitempty"`
	Code    string `json:"code"`
	Status  int    `json:"-"`
}

func (e ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.Status)
	return nil
}

func (s GenericResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
