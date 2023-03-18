package news

import (
	"context"

	"github.com/friendsofgo/errors"

	"github.com/flowck/news_service_ddd_golang/internal/domain"
)

var (
	ErrNewsNotFound = errors.New("news not found")
)

type Repository interface {
	Insert(ctx context.Context, n *News, topicIDs []domain.ID) error
}
