package commands

import (
	"context"

	"github.com/flowck/news_service_ddd_golang/internal/domain"

	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
)

type PublishNews struct {
	News     *news.News
	TopicIds []domain.ID
}

type publishNewsHandler struct{}

func NewPublishNewsHandler() publishNewsHandler {
	return publishNewsHandler{}
}

func (p publishNewsHandler) Execute(ctx context.Context, cmd PublishNews) error {
	return nil
}
