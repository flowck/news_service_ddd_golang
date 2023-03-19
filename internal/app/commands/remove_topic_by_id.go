package commands

import (
	"context"

	"github.com/flowck/news_service_ddd_golang/internal/domain"
	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
)

type RemoveTopicByID struct {
	ID domain.ID
}

type removeTopicByIdHandler struct {
	repository news.TopicRepository
}

func NewRemoveTopicByIDHandler(repository news.TopicRepository) removeTopicByIdHandler {
	return removeTopicByIdHandler{repository: repository}
}

func (p removeTopicByIdHandler) Execute(ctx context.Context, cmd RemoveTopicByID) error {
	return p.repository.DeleteByID(ctx, cmd.ID)
}
