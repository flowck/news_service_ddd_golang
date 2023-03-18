package queries

import (
	"context"

	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
)

type NewsReadModel interface {
	FindNewsByFiltersWithPagination(ctx context.Context, f NewsFilter) (NewsPaginated, error)
}

type NewsFilter struct {
	Limit int
	Page  int

	Status string
	Topic  string
}

type NewsPaginated struct {
	Page         int
	TotalPages   int64
	TotalResults int64
	Data         []*news.News
}
