package newsarticles

import (
	"github.com/flowck/news_service_ddd_golang/internal/adapters/models"
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
