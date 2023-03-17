package news

import (
	"errors"
	"time"

	"github.com/flowck/news_service_ddd_golang/internal/domain"
)

type News struct {
	id           domain.ID
	title        string
	slug         Slug
	content      string
	status       Status
	topics       []*Topic
	publishedAt  time.Time
	lastEditedAt time.Time
}

func (n *News) Id() domain.ID {
	return n.id
}

func (n *News) Title() string {
	return n.title
}

func (n *News) Slug() Slug {
	return n.slug
}

func (n *News) Content() string {
	return n.content
}

func (n *News) Status() Status {
	return n.status
}

func (n *News) Topics() []*Topic {
	return n.topics
}

func (n *News) PublishedAt() time.Time {
	return n.publishedAt
}

func (n *News) LastEditedAt() time.Time {
	return n.lastEditedAt
}

func New(
	id domain.ID,
	title string,
	slug Slug,
	content string,
	status Status,
	topics []*Topic,
	publishedAt,
	lastEditedAt time.Time,
) (*News, error) {
	if id.IsZero() {
		return nil, errors.New("id cannot be invalid")
	}

	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	if content == "" {
		return nil, errors.New("content cannot be empty")
	}

	if slug.IsEmpty() {
		return nil, errors.New("slug cannot be invalid")
	}

	if status.IsZero() {
		return nil, errors.New("status cannot be invalid")
	}

	if publishedAt.IsZero() {
		return nil, errors.New("publishedAt cannot be invalid")
	}

	if lastEditedAt.IsZero() {
		return nil, errors.New("lastEditedAt cannot be invalid")
	}

	return &News{
		id:           id,
		title:        title,
		slug:         slug,
		content:      content,
		status:       status,
		topics:       topics,
		publishedAt:  publishedAt,
		lastEditedAt: lastEditedAt,
	}, nil
}
