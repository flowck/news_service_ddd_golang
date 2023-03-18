package transaction

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/flowck/news_service_ddd_golang/internal/adapters/newsarticles"
	"github.com/flowck/news_service_ddd_golang/internal/app/commands"
	"github.com/flowck/news_service_ddd_golang/internal/common/logs"
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

	newsRepository, err := newsarticles.NewPsqlRepository(tx)
	if err != nil {
		mustRollbackTx(tx, err)
		return err
	}

	adapters := commands.TransactableAdapters{
		NewsRepository: newsRepository,
	}

	err = f(adapters)
	if err != nil {
		mustRollbackTx(tx, err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func mustRollbackTx(tx *sql.Tx, err error) {
	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		// p.logger.WithError(err).WithField("rollback_err", rollbackErr).Error("Rollback error")
		logs.Errorf("tx failed due to: %v", err)
		logs.Errorf("rollback error: %v", rollbackErr)
	}
}
