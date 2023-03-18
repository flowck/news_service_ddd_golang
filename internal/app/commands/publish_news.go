package commands

import (
	"context"

	"github.com/flowck/news_service_ddd_golang/internal/domain"
	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
)

type PublishNews struct {
	News     *news.News
	TopicIds []domain.ID
}

type publishNewsHandler struct {
	txProvider TransactionProvider
}

func NewPublishNewsHandler(txProvider TransactionProvider) publishNewsHandler {
	return publishNewsHandler{txProvider: txProvider}
}

func (p publishNewsHandler) Execute(ctx context.Context, cmd PublishNews) error {
	return p.txProvider.Transact(ctx, func(adapters TransactableAdapters) error {
		err := adapters.NewsRepository.Insert(ctx, cmd.News, cmd.TopicIds)
		if err != nil {
			return err
		}

		return nil
	})
}
