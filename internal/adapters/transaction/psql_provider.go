package transaction

import (
	"context"

	"github.com/flowck/news_service_ddd_golang/internal/common/logs"

	"github.com/flowck/news_service_ddd_golang/internal/app/commands"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type SQLProvider struct {
	db boil.ContextBeginner
}

func NewSQLProvider(db boil.ContextBeginner) SQLProvider {
	return SQLProvider{
		db: db,
	}
}

func (p SQLProvider) Transact(ctx context.Context, f commands.TransactFunc) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	adapters := commands.TransactableAdapters{}

	err = f(adapters)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			// p.logger.WithError(err).WithField("rollback_err", rollbackErr).Error("Rollback error")
			logs.Error("rollback error")
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
