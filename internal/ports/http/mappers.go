package http

import (
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
