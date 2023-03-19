package queries

import (
	"context"

	"github.com/flowck/news_service_ddd_golang/internal/domain"
	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
)

type TopicByID struct {
	ID domain.ID
}

type topicByIdHandler struct {
	repository news.TopicRepository
}

func NewTopicByIdHandler(repository news.TopicRepository) topicByIdHandler {
	return topicByIdHandler{repository: repository}
}

func (h topicByIdHandler) Execute(ctx context.Context, q TopicByID) (*news.Topic, error) {
	return h.repository.Find(ctx, q.ID)
}
