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

func reply(w http.ResponseWriter, r *http.Request, renderer render.Renderer) {
	if err := render.Render(w, r, renderer); err != nil {
		render.Respond(w, r, newInternalServerError(err))
	}
}
