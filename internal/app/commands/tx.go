package commands

import "context"

type TransactionProvider interface {
	Transact(ctx context.Context, f TransactFunc) error
}

type TransactableAdapters struct {
}

type TransactFunc func(adapters TransactableAdapters) error
