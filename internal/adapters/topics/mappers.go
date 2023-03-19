package topics

import (
	"github.com/flowck/news_service_ddd_golang/internal/adapters/models"
	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
)

func mapTopicDomainToTopicModel(topic *news.Topic) *models.Topic {
	return &models.Topic{
		ID:   topic.ID().String(),
		Name: topic.Value(),
	}
}
