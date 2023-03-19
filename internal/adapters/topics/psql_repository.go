package topics

import (
	"context"
	"database/sql"

	"github.com/flowck/news_service_ddd_golang/internal/domain"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/flowck/news_service_ddd_golang/internal/adapters/models"
	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
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

func (p psqlRepository) FindAll(ctx context.Context) ([]*news.Topic, error) {
	rows, err := models.Topics().All(ctx, p.db)
	if err != nil {
		return nil, err
	}

	return mapTopicListModelToTopicListDomain(rows)
}

func (p psqlRepository) Find(ctx context.Context, ID domain.ID) (*news.Topic, error) {
	row, err := models.FindTopic(ctx, p.db, ID.String())
	if errors.Is(err, sql.ErrNoRows) {
		return nil, news.ErrTopicNotFound
	}

	return mapTopicModelToTopicDomain(row)
}

func (p psqlRepository) DeleteByID(ctx context.Context, ID domain.ID) error {
	row, err := models.FindTopic(ctx, p.db, ID.String())
	if errors.Is(err, sql.ErrNoRows) {
		return news.ErrTopicNotFound
	}

	if err != nil {
		return err
	}

	if _, err = row.Delete(ctx, p.db); err != nil {
		return err
	}

	return nil
}
