package commands

import (
	"context"
	"time"

	"github.com/flowck/news_service_ddd_golang/internal/domain"
	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
)

type EditNews struct {
	ID domain.ID

	Content     string
	Slug        news.Slug
	Status      news.Status
	Title       string
	PublishedAt time.Time
	TopicsIds   []domain.ID
}

type editNewsHandler struct {
	repository news.Repository
}

func NewEditNewsHandler(repository news.Repository) editNewsHandler {
	return editNewsHandler{
		repository: repository,
	}
}

func (p editNewsHandler) Execute(ctx context.Context, cmd EditNews) error {
	err := p.repository.Update(ctx, cmd.ID, cmd.TopicsIds, func(n *news.News) error {
		err := n.Edit(cmd.Title, cmd.Content, cmd.Slug, cmd.Status, cmd.PublishedAt)
		if err != nil {
			return nil
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
