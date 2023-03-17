package http

import (
	"github.com/flowck/news_service_ddd_golang/internal/ports/http/static"
	nethttp "net/http"
)

func (h handlers) GetNews(w nethttp.ResponseWriter, r *nethttp.Request, params static.GetNewsParams) {
	reply(w, r, &GenericResponse{"success"})
}
