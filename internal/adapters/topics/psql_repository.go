package topics

import (
	"context"

	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type psqlRepository struct {
	db boil.ContextExecutor
}

func NewPsqlRepository(db boil.ContextExecutor) (psqlRepository, error) {
	if db == nil {
		return psqlRepository{}, errors.New("db cannot be nil")
	}

	return psqlRepository{db: db}, nil
}

func (p psqlRepository) Insert(ctx context.Context, t *news.Topic) error {
	row := mapTopicDomainToTopicModel(t)

	if err := row.Insert(ctx, p.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}
