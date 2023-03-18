package newsarticles

import (
	"context"
	"errors"

	"github.com/flowck/news_service_ddd_golang/internal/adapters/models"
	"github.com/flowck/news_service_ddd_golang/internal/domain"
	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ news.Repository = (*psqlRepository)(nil)

type psqlRepository struct {
	db boil.ContextExecutor
}

func NewPsqlRepository(db boil.ContextExecutor) (psqlRepository, error) {
	if db == nil {
		return psqlRepository{}, errors.New("db cannot be nil")
	}

	return psqlRepository{db: db}, nil
}

func (p psqlRepository) Insert(ctx context.Context, n *news.News, topicIDs []domain.ID) error {
	row := mapNewsToDbModel(n)

	if err := row.Insert(ctx, p.db, boil.Infer()); err != nil {
		return err
	}

	if topicIDs != nil {
		err := p.saveNewsArticleTopics(ctx, n.Id(), topicIDs)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p psqlRepository) saveNewsArticleTopics(
	ctx context.Context,
	nID domain.ID,
	topicIDs []domain.ID,
) error {
	for i := range topicIDs {
		row := models.NewsArticlesTopic{
			TopicID:        topicIDs[i].String(),
			NewsArticlesID: nID.String(),
		}

		if err := row.Insert(ctx, p.db, boil.Infer()); err != nil {
			return err
		}
	}

	return nil
}
