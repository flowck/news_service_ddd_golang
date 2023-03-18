package http

import (
	"fmt"
	nethttp "net/http"
	"strconv"

	"github.com/flowck/news_service_ddd_golang/internal/domain"
	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
	"github.com/flowck/news_service_ddd_golang/internal/ports/http/static"
)

func mapTopicIdsToDomainIds(topicIds []string) ([]domain.ID, error) {
	IDs := make([]domain.ID, len(topicIds))

	for i, id := range topicIds {
		ID, err := domain.NewIDFromString(id)
		if err != nil {
			return nil, err
		}

		IDs[i] = ID
	}

	return IDs, nil
}

func mapDomainNewsToResponseNews(n *news.News) static.News {
	topics := make([]string, 1)

	if n.Topics() != nil {
		topics = mapDomainTopicsToResponse(n.Topics())
	}

	return static.News{
		Id:           n.Id().String(),
		Title:        n.Title(),
		Slug:         n.Slug().String(),
		Status:       n.Status().String(),
		Content:      n.Content(),
		Topics:       topics,
		LastEditedAt: n.LastEditedAt(),
		PublishedAt:  n.PublishedAt(),
	}
}

func mapDomainTopicsToResponse(topics []*news.Topic) []string {
	result := make([]string, len(topics))

	for i, topic := range topics {
		result[i] = topic.String()
	}

	return result
}

func toPtr[V any](v V) *V {
	return &v
}

func mapRequestQueryToPaginationParams(r *nethttp.Request) (static.GetNewsParams, error) {
	l := r.URL.Query().Get("limit")
	limit := 100
	var err error

	if l != "" {
		if limit, err = strconv.Atoi(l); err != nil {
			return static.GetNewsParams{}, fmt.Errorf("limit '%s' is not a valid integer", l)
		}

		if limit > 100 {
			return static.GetNewsParams{}, fmt.Errorf("limit '%s' is greater than 100", l)
		}
	}

	page := 1
	p := r.URL.Query().Get("page")

	if p != "" {
		if page, err = strconv.Atoi(p); err != nil {
			return static.GetNewsParams{}, fmt.Errorf("page '%s' is not a valid integer", p)
		}
	}

	return static.GetNewsParams{
		Limit:  toPtr(limit),
		Page:   toPtr(page),
		Status: toPtr(r.URL.Query().Get("status")),
		Topic:  toPtr(r.URL.Query().Get("topic")),
	}, nil
}
