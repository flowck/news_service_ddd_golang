package transaction

import (
	"context"

	"github.com/flowck/news_service_ddd_golang/internal/app/commands"
)

var _ commands.TransactionProvider = (*psqlProviderMock)(nil)

type psqlProviderMock struct {
}

func (p psqlProviderMock) Transact(ctx context.Context, f commands.TransactFunc) error {
	adapters := commands.TransactableAdapters{}

	if err := f(adapters); err != nil {
		return err
	}

	return nil
}

func NewPsqlProviderMock() (psqlProviderMock, error) {
	return psqlProviderMock{}, nil
}
