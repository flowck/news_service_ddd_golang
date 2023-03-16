package http

import nethttp "net/http"

func (h handlers) GetNews(w nethttp.ResponseWriter, r *nethttp.Request) {
	reply(w, r, &GenericResponse{"success"})
}
