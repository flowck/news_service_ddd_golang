package http

import (
	"net/http"

	"github.com/go-chi/render"
)

func newInternalServerError(err error) *ErrResponse {
	return &ErrResponse{
		Err:     err,
		Message: "Error",
		Code:    "internal_server_error",
		Status:  http.StatusInternalServerError,
	}
}

func newErrorWithStatus(err error, code string, status int) *ErrResponse {
	return &ErrResponse{
		Err:     err,
		Message: err.Error(),
		Code:    code,
		Status:  status,
	}
}

func reply(w http.ResponseWriter, r *http.Request, renderer render.Renderer) {
	if err := render.Render(w, r, renderer); err != nil {
		render.Respond(w, r, newInternalServerError(err))
	}
}

func replyWithStatus(w http.ResponseWriter, r *http.Request, renderer render.Renderer, status int) {
	w.WriteHeader(status)
	reply(w, r, renderer)
}

func bind(w http.ResponseWriter, r *http.Request, binder render.Binder) error {
	if err := render.Bind(r, binder); err != nil {
		render.Respond(w, r, ErrResponse{
			Err:     err,
			Message: "Unable to parse the request body",
			Code:    "error_parsing_body",
			Status:  http.StatusUnprocessableEntity,
		})
		return err
	}

	return nil
}
