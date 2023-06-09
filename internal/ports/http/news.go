package http

import (
	"errors"
	nethttp "net/http"

	"github.com/flowck/news_service_ddd_golang/internal/app/commands"
	"github.com/flowck/news_service_ddd_golang/internal/app/queries"
	"github.com/flowck/news_service_ddd_golang/internal/domain"
	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
	"github.com/flowck/news_service_ddd_golang/internal/ports/http/static"
)

func (h handlers) GetNews(w nethttp.ResponseWriter, r *nethttp.Request, params static.GetNewsParams) {
	query := queries.NewsByFilters{Filters: queries.NewsFilter{
		Limit:  *params.Limit,
		Page:   *params.Page,
		Status: *params.Status,
		Topic:  *params.Topic,
	}}

	paginatedNews, err := h.application.Queries.NewsByFilters.Execute(r.Context(), query)
	if err != nil {
		reply(w, r, newErrorWithStatus(err, "unable-to-get-news", nethttp.StatusInternalServerError))
		return
	}

	reply(w, r, &GetNewsByFiltersPayload{
		Data:         mapDomainNewsListToResponseNewsList(paginatedNews.Data),
		Page:         paginatedNews.Page,
		TotalPages:   paginatedNews.TotalPages,
		TotalResults: paginatedNews.TotalResults,
	})
}

func (h handlers) GetNewsByID(w nethttp.ResponseWriter, r *nethttp.Request, newsID string) {
	ID, err := domain.NewIDFromString(newsID)
	if err != nil {
		reply(w, r, newErrorWithStatus(err, "invalid-news-id", nethttp.StatusBadRequest))
		return
	}

	n, err := h.application.Queries.NewsByID.Execute(r.Context(), queries.NewsByID{ID: ID})
	if errors.Is(err, news.ErrNewsNotFound) {
		reply(w, r, newErrorWithStatus(err, "news-not-found", nethttp.StatusNotFound))
		return
	}

	reply(w, r, &GetNewsPayload{
		Data: mapDomainNewsToResponseNews(n),
	})
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

	replyWithStatus(w, r, GenericResponse{Message: "news created successfully"}, nethttp.StatusCreated)
}

func (h handlers) UnPublishNews(w nethttp.ResponseWriter, r *nethttp.Request, newsID string) {
	ID, err := domain.NewIDFromString(newsID)
	if err != nil {
		reply(w, r, newErrorWithStatus(err, "invalid-news-id", nethttp.StatusBadRequest))
		return
	}

	err = h.application.Commands.UnPublishNews.Execute(r.Context(), commands.UnPublishNews{
		ID: ID,
	})

	if errors.Is(err, news.ErrNewsNotFound) {
		reply(w, r, newErrorWithStatus(err, "news-not-found", nethttp.StatusNotFound))
		return
	}

	if err != nil {
		reply(w, r, newErrorWithStatus(err, "unable-to-unpublish-news", nethttp.StatusInternalServerError))
		return
	}

	replyWithStatus(w, r, GenericResponse{}, nethttp.StatusNoContent)
}

func (h handlers) EditNews(w nethttp.ResponseWriter, r *nethttp.Request, newsID string) {
	ID, err := domain.NewIDFromString(newsID)
	if err != nil {
		reply(w, r, newErrorWithStatus(err, "invalid-news-id", nethttp.StatusBadRequest))
		return
	}

	body := &EditNewsRequest{}
	if err = bind(w, r, body); err != nil {
		return
	}

	sl, err := news.NewSlugFromString(body.Slug)
	if err != nil {
		reply(w, r, newErrorWithStatus(err, "invalid-slug", nethttp.StatusBadRequest))
		return
	}

	status, err := news.NewStatusFromString(body.Status)
	if err != nil {
		reply(w, r, newErrorWithStatus(err, "invalid-status", nethttp.StatusBadRequest))
		return
	}

	cmd := commands.EditNews{
		ID:          ID,
		Content:     body.Content,
		Slug:        sl,
		Status:      status,
		Title:       body.Title,
		PublishedAt: body.PublishedAt,
		TopicsIds:   nil,
	}
	err = h.application.Commands.EditNews.Execute(r.Context(), cmd)
	if errors.Is(err, news.ErrNewsNotFound) {
		reply(w, r, newErrorWithStatus(err, "news-not-found", nethttp.StatusNotFound))
		return
	}

	if err != nil {
		reply(w, r, newErrorWithStatus(err, "unable-to-edit-news", nethttp.StatusInternalServerError))
		return
	}

	replyWithStatus(w, r, nil, nethttp.StatusNoContent)
}
