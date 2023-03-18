package commands

import (
	"context"

	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
)

type TransactionProvider interface {
	Transact(ctx context.Context, f TransactFunc) error
}

type TransactableAdapters struct {
	NewsRepository news.Repository
}

type TransactFunc func(adapters TransactableAdapters) error
