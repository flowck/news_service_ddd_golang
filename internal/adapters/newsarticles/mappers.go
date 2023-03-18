package newsarticles

import (
	"github.com/flowck/news_service_ddd_golang/internal/adapters/models"
	"github.com/flowck/news_service_ddd_golang/internal/domain"
	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
	"github.com/volatiletech/null/v8"
)

func mapNewsToDbModel(n *news.News) *models.NewsArticle {
	return &models.NewsArticle{
		ID:           n.Id().String(),
		Title:        n.Title(),
		Slug:         n.Slug().String(),
		Content:      null.StringFrom(n.Content()),
		Status:       null.StringFrom(n.Status().String()),
		PublishedAt:  null.TimeFrom(n.PublishedAt()),
		LastEditedAt: null.TimeFrom(n.LastEditedAt()),
	}
}

func mapNewsModelToNewsDomain(row *models.NewsArticle) (*news.News, error) {
	ID, err := domain.NewIDFromString(row.ID)
	if err != nil {
		return nil, err
	}

	sl, err := news.NewSlugFromString(row.Slug)
	if err != nil {
		return nil, err
	}

	status, err := news.NewStatusFromString(row.Status.String)
	if err != nil {
		return nil, err
	}

	n, err := news.New(
		ID,
		row.Title,
		sl,
		row.Content.String,
		status,
		nil,
		row.PublishedAt.Time,
		row.LastEditedAt.Time,
	)
	if err != nil {
		return nil, err
	}

	return n, nil
}
