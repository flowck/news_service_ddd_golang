package topics

import (
	"github.com/flowck/news_service_ddd_golang/internal/adapters/models"
	"github.com/flowck/news_service_ddd_golang/internal/domain"
	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
)

func mapTopicDomainToTopicModel(topic *news.Topic) *models.Topic {
	return &models.Topic{
		ID:   topic.ID().String(),
		Name: topic.Value(),
	}
}

func mapTopicModelToTopicDomain(row *models.Topic) (*news.Topic, error) {
	ID, err := domain.NewIDFromString(row.ID)
	if err != nil {
		return nil, err
	}

	return news.NewTopic(ID, row.Name)
}

func mapTopicListModelToTopicListDomain(rows []*models.Topic) ([]*news.Topic, error) {
	result := make([]*news.Topic, len(rows))

	for i, row := range rows {
		t, err := mapTopicModelToTopicDomain(row)
		if err != nil {
			return nil, err
		}

		result[i] = t
	}

	return result, nil
}
