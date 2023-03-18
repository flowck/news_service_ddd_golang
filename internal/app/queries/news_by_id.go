package queries

import (
	"context"

	"github.com/flowck/news_service_ddd_golang/internal/domain"
	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
)

type NewsByID struct {
	ID domain.ID
}

type newsByIdHandler struct {
	repository news.Repository
}

func NewNewsByIdHandler(repository news.Repository) newsByIdHandler {
	return newsByIdHandler{repository: repository}
}

func (h newsByIdHandler) Execute(ctx context.Context, q NewsByID) (*news.News, error) {
	return h.repository.Find(ctx, q.ID)
}
