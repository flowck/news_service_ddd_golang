package newsarticles

import (
	"context"
	"database/sql"
	"errors"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/flowck/news_service_ddd_golang/internal/adapters/models"
	"github.com/flowck/news_service_ddd_golang/internal/app/queries"
	"github.com/flowck/news_service_ddd_golang/internal/domain"
	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
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

func (p psqlRepository) Find(ctx context.Context, ID domain.ID) (*news.News, error) {
	mods := []qm.QueryMod{
		models.NewsArticleWhere.ID.EQ(ID.String()),
		qm.Load(models.NewsArticleRels.Topics),
	}
	row, err := models.NewsArticles(mods...).One(ctx, p.db)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, news.ErrNewsNotFound
	}

	if err != nil {
		return nil, err
	}

	return mapNewsModelToNewsDomain(row)
}

func (p psqlRepository) Update(ctx context.Context, n *news.News) error {
	row := mapNewsToDbModel(n)

	if _, err := row.Update(ctx, p.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (p psqlRepository) Insert(ctx context.Context, n *news.News, topicIDs []domain.ID) error {
	row := mapNewsToDbModel(n)

	if err := row.Insert(ctx, p.db, boil.Infer()); err != nil {
		return err
	}

	if topicIDs != nil {
		err := p.setTopicsToNewsArticle(ctx, row, topicIDs)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p psqlRepository) FindNewsByFiltersWithPagination(
	ctx context.Context,
	f queries.NewsFilter,
) (queries.NewsPaginated, error) {
	mods := []qm.QueryMod{
		qm.Offset(getPaginationOffset(f.Page, f.Limit)),
		qm.Load(models.NewsArticleRels.Topics, qm.Where("name ILIKE %?", f.Topic)),
	}

	if f.Status != "" {
		mods = append(mods, models.NewsArticleWhere.Status.EQ(null.StringFrom(f.Status)))
	}

	modsWithLimit := append(mods, qm.Limit(f.Limit))
	rows, err := models.NewsArticles(modsWithLimit...).All(ctx, p.db)
	if err != nil {
		return queries.NewsPaginated{}, err
	}

	newsList, err := mapNewsListModelToNewsListDomain(rows)
	if err != nil {
		return queries.NewsPaginated{}, err
	}

	totalRows, err := models.NewsArticles(mods...).Count(ctx, p.db)
	if err != nil {
		return queries.NewsPaginated{}, err
	}

	return queries.NewsPaginated{
		Page:         f.Page,
		TotalPages:   getTotalPages(totalRows, f.Limit),
		TotalResults: totalRows,
		Data:         newsList,
	}, nil
}

func getTotalPages(totalRows int64, limit int) int64 {
	return totalRows / int64(limit)
}

func getPaginationOffset(page int, limit int) int {
	return (page - 1) * limit
}

func (p psqlRepository) setTopicsToNewsArticle(
	ctx context.Context,
	newsArticleRow *models.NewsArticle,
	topicIDs []domain.ID,
) error {
	rows := make([]*models.Topic, len(topicIDs))

	for i := range topicIDs {
		row, err := models.FindTopic(ctx, p.db, topicIDs[i].String())
		if err != nil {
			return err
		}

		rows[i] = row
	}

	if err := newsArticleRow.SetTopics(ctx, p.db, false, rows...); err != nil {
		return err
	}

	return nil
}
