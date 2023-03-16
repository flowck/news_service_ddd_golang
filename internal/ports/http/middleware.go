package http

import (
	"encoding/json"
	"fmt"
	nethttp "net/http"

	"github.com/flowck/news_service_ddd_golang/internal/common/logs"
	"github.com/go-chi/cors"
)

func corsConfig(allowedOrigin []string) cors.Options {
	return cors.Options{
		AllowedOrigins:   allowedOrigin,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}
}

func globalErrorHandler(w nethttp.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	res, err := json.Marshal(&ErrResponse{
		Message: message,
		Code:    "general_error",
	})

	if err != nil {
		logs.Errorf("unable to marshal error message: %v", err)
		fmt.Fprint(w, `{"message": "internal-server-error"}`)
		return
	}

	fmt.Fprint(w, string(res))
}
