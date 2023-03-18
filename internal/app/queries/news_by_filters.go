package queries

import (
	"context"

	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
)

type NewsByFilters struct {
	Filters NewsFilter
}

type newsByFiltersHandler struct {
	readModelRepository NewsReadModel
}

func NewNewsByFiltersHandler(readModelRepository NewsReadModel) newsByFiltersHandler {
	return newsByFiltersHandler{readModelRepository: readModelRepository}
}

func (h newsByFiltersHandler) Execute(ctx context.Context, q NewsByFilters) ([]*news.News, error) {
	return h.readModelRepository.FindNewsByFiltersWithPagination(ctx, q.Filters)
}
