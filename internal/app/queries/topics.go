package queries

import (
	"context"

	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
)

type Topics struct {
}

type topicsHandler struct {
	repository news.TopicRepository
}

func NewTopicsHandler(repository news.TopicRepository) topicsHandler {
	return topicsHandler{repository: repository}
}

func (h topicsHandler) Execute(ctx context.Context, q Topics) ([]*news.Topic, error) {
	return h.repository.Find(ctx)
}
