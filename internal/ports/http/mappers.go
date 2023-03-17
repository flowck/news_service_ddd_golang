package http

import "github.com/flowck/news_service_ddd_golang/internal/domain"

func mapTopicIdsToDomainIds(topicIds []string) ([]domain.ID, error) {
	IDs := make([]domain.ID, len(topicIds))

	for i, id := range topicIds {
		ID, err := domain.NewIDFromString(id)
		if err != nil {
			return nil, err
		}

		IDs[i] = ID
	}

	return IDs, nil
}
