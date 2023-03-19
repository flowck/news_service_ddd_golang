package http

import (
	nethttp "net/http"

	"github.com/flowck/news_service_ddd_golang/internal/app/commands"
	"github.com/flowck/news_service_ddd_golang/internal/app/queries"
	"github.com/flowck/news_service_ddd_golang/internal/domain"
	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
)

func (h handlers) CreateTopic(w nethttp.ResponseWriter, r *nethttp.Request) {
	body := &CreateTopicRequest{}
	if err := bind(w, r, body); err != nil {
		return
	}

	topic, err := news.NewTopic(domain.NewID(), body.Name)
	if err != nil {
		reply(w, r, newErrorWithStatus(err, "invalid-topic", nethttp.StatusBadRequest))
		return
	}

	cmd := commands.CreateTopic{Topic: topic}
	err = h.application.Commands.CreateTopic.Execute(r.Context(), cmd)
	if err != nil {
		reply(w, r, newErrorWithStatus(err, "unable-to-create-topic", nethttp.StatusInternalServerError))
		return
	}

	replyWithStatus(w, r, GenericResponse{Message: "topic created successfully"}, nethttp.StatusCreated)
}

func (h handlers) GetTopics(w nethttp.ResponseWriter, r *nethttp.Request) {
	topics, err := h.application.Queries.Topics.Execute(r.Context(), queries.Topics{})
	if err != nil {
		reply(w, r, newErrorWithStatus(err, "unable-to-get-topics", nethttp.StatusInternalServerError))
		return
	}

	reply(w, r, GetTopicsPayload{Data: mapDomainTopicsToResponse(topics)})
}
