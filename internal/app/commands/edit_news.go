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
	Slug        string
	Status      string
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
	return nil
}
