package commands

import (
	"context"

	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
)

type CreateTopic struct {
	Topic news.Topic
}

type createTopicHandler struct {
}

func NewCreateTopicHandler() createTopicHandler {
	return createTopicHandler{}
}

func (p createTopicHandler) Execute(ctx context.Context, cmd CreateTopic) error {
	return nil
}
