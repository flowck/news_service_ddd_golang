package http

import (
	"net/http"

	"github.com/flowck/news_service_ddd_golang/internal/ports/http/static"
)

type GenericResponse struct {
	Message string `json:"message,omitempty"`
}

type PublishNewsRequest static.PublishNewsRequest
type GetNewsPayload static.GetNewsPayload
type GetNewsByFiltersPayload static.GetNewsByFiltersPayload
type CreateTopicRequest static.CreateTopicRequest
type GetTopicsPayload static.GetTopicsPayload
type GetTopicPayload static.GetTopicPayload
type EditTopicRequest static.EditTopicRequest

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

func (s GetNewsByFiltersPayload) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s GenericResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s GetNewsPayload) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s GetTopicsPayload) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s GetTopicPayload) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (p PublishNewsRequest) Bind(r *http.Request) error {
	return nil
}

func (p CreateTopicRequest) Bind(r *http.Request) error {
	return nil
}

func (p EditTopicRequest) Bind(r *http.Request) error {
	return nil
}
