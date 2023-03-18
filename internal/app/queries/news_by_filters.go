package queries

import (
	"context"
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

func (h newsByFiltersHandler) Execute(ctx context.Context, q NewsByFilters) (NewsPaginated, error) {
	return h.readModelRepository.FindNewsByFiltersWithPagination(ctx, q.Filters)
}
