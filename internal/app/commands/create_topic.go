package commands

import (
	"context"

	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
)

type CreateTopic struct {
	Topic *news.Topic
}

type createTopicHandler struct {
	repository news.TopicRepository
}

func NewCreateTopicHandler(repository news.TopicRepository) createTopicHandler {
	return createTopicHandler{
		repository: repository,
	}
}

func (p createTopicHandler) Execute(ctx context.Context, cmd CreateTopic) error {
	return p.repository.Insert(ctx, cmd.Topic)
}
