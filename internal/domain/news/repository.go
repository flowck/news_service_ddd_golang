package news

import (
	"context"

	"github.com/friendsofgo/errors"

	"github.com/flowck/news_service_ddd_golang/internal/domain"
)

var (
	ErrNewsNotFound  = errors.New("news not found")
	ErrTopicNotFound = errors.New("topic not found")
)

type Repository interface {
	Insert(ctx context.Context, n *News, topicIDs []domain.ID) error
	Update(ctx context.Context, n *News) error
	Find(ctx context.Context, ID domain.ID) (*News, error)
}

type TopicRepository interface {
	Insert(ctx context.Context, t *Topic) error
	Find(ctx context.Context, ID domain.ID) (*Topic, error)
	FindAll(ctx context.Context) ([]*Topic, error)
	DeleteByID(ctx context.Context, ID domain.ID) error
	Update(ctx context.Context, topic *Topic) error
}
