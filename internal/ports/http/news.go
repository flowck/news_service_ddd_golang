package http

import (
	nethttp "net/http"

	"github.com/flowck/news_service_ddd_golang/internal/app/commands"
	"github.com/flowck/news_service_ddd_golang/internal/domain"
	"github.com/flowck/news_service_ddd_golang/internal/domain/news"

	"github.com/flowck/news_service_ddd_golang/internal/ports/http/static"
)

func (h handlers) GetNews(w nethttp.ResponseWriter, r *nethttp.Request, params static.GetNewsParams) {
	reply(w, r, &GenericResponse{"success"})
}

func (h handlers) PublishNews(w nethttp.ResponseWriter, r *nethttp.Request) {
	body := &PublishNewsRequest{}
	if err := bind(w, r, body); err != nil {
		return
	}

	newsID := domain.NewID()

	sl, err := news.NewSlugFromID(newsID, body.Slug)
	if err != nil {
		reply(w, r, newErrorWithStatus(err, "invalid-slug", nethttp.StatusBadRequest))
		return
	}

	status, err := news.NewStatusFromString(body.Status)
	if err != nil {
		reply(w, r, newErrorWithStatus(err, "invalid-status", nethttp.StatusBadRequest))
		return
	}

	n, err := news.New(
		domain.NewID(),
		body.Title,
		sl,
		body.Content,
		status,
		nil,
		body.PublishedAt,
		body.PublishedAt,
	)

	if err != nil {
		reply(w, r, newErrorWithStatus(err, "invalid-news", nethttp.StatusBadRequest))
		return
	}

	topicIDs, err := mapTopicIdsToDomainIds(body.TopicIds)
	if err != nil {
		reply(w, r, newErrorWithStatus(err, "invalid-topic-id", nethttp.StatusBadRequest))
		return
	}

	err = h.application.Commands.PublishNews.Execute(r.Context(), commands.PublishNews{
		News:     n,
		TopicIds: topicIDs,
	})
	if err != nil {
		reply(w, r, newErrorWithStatus(err, "unable-to-publish", nethttp.StatusInternalServerError))
		return
	}

	replyWithStatus(w, r, GenericResponse{}, nethttp.StatusCreated)
}
