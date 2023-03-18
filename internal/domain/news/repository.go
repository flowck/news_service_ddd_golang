package news

import (
	"context"

	"github.com/flowck/news_service_ddd_golang/internal/domain"
)

type Repository interface {
	Insert(ctx context.Context, n *News, topicIDs []domain.ID) error
}
