package commands

import (
	"context"

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

func (p unPublishNewsHandler) Execute(ctx context.Context, cmd PublishNews) error {
	return nil
}
