package commands

import (
	"context"

	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
)

type EditTopic struct {
	Topic *news.Topic
}

type editTopicHandler struct {
	repository news.TopicRepository
}

func NewEditTopicHandler(repository news.TopicRepository) editTopicHandler {
	return editTopicHandler{repository: repository}
}

func (p editTopicHandler) Execute(ctx context.Context, cmd EditTopic) error {
	return p.repository.Update(ctx, cmd.Topic)
}
