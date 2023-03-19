package news

import (
	"github.com/flowck/news_service_ddd_golang/internal/domain"
	"github.com/friendsofgo/errors"
)

type Topic struct {
	id    domain.ID
	value string
}

func NewTopic(ID domain.ID, value string) (*Topic, error) {
	if ID.IsZero() {
		return nil, errors.New("id cannot be invalid")
	}

	if value == "" {
		return nil, errors.New("value cannot be empty")
	}

	return &Topic{
		id:    ID,
		value: value,
	}, nil
}

func (t Topic) String() string {
	return t.value
}

func (t Topic) ID() domain.ID {
	return t.id
}

func (t Topic) Value() string {
	return t.value
}
