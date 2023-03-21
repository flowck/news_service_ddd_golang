package commands

import (
	"context"

	"github.com/flowck/news_service_ddd_golang/internal/domain/news"

	"github.com/flowck/news_service_ddd_golang/internal/domain"
)

type UnPublishNews struct {
	ID domain.ID
}

type unPublishNewsHandler struct {
	txProvider TransactionProvider
}

func NewUnPublishNewsHandler(txProvider TransactionProvider) unPublishNewsHandler {
	return unPublishNewsHandler{txProvider: txProvider}
}

func (p unPublishNewsHandler) Execute(ctx context.Context, cmd UnPublishNews) error {
	return p.txProvider.Transact(ctx, func(adapters TransactableAdapters) error {
		err := adapters.NewsRepository.Update(ctx, cmd.ID, nil, func(n *news.News) error {
			n.UnPublish()
			return nil
		})

		if err != nil {
			return nil
		}

		return nil
	})
}
