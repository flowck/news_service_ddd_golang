package queries

import (
	"context"

	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
)

type NewsReadModel interface {
	FindNewsByFiltersWithPagination(ctx context.Context, f NewsFilter) ([]*news.News, error)
}

type NewsFilter struct {
	Limit int
	Page  int

	Status string
	Topic  string
}
