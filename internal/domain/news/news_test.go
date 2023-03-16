package news_test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/flowck/news_service_ddd_golang/internal/domain"
	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
	"github.com/gosimple/slug"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	id01 := domain.MustNewIDFromString(gofakeit.UUID())

	testCases := []struct {
		name        string
		expectedErr string

		id           domain.ID
		title        string
		slug         news.Slug
		content      string
		status       news.Status
		topics       []*news.Topic
		publishedAt  time.Time
		lastEditedAt time.Time
	}{
		{
			name:        "invalid_title",
			expectedErr: "id cannot be invalid",

			id:           domain.ID{},
			title:        gofakeit.Sentence(40),
			slug:         news.MustNewSlugFromID(id01, slug.Make(gofakeit.Quote())),
			content:      gofakeit.Sentence(500),
			status:       news.StatusPublished,
			topics:       nil,
			publishedAt:  time.Now(),
			lastEditedAt: time.Now(),
		},
		{
			name:        "empty_title",
			expectedErr: "title cannot be empty",

			id:           id01,
			title:        "",
			slug:         news.MustNewSlugFromID(id01, slug.Make(gofakeit.Quote())),
			content:      gofakeit.Sentence(500),
			status:       news.StatusPublished,
			topics:       nil,
			publishedAt:  time.Now(),
			lastEditedAt: time.Now(),
		},
		{
			name:        "empty_content",
			expectedErr: "content cannot be empty",

			id:           id01,
			title:        gofakeit.Sentence(40),
			slug:         news.MustNewSlugFromID(id01, slug.Make(gofakeit.Quote())),
			content:      "",
			status:       news.StatusPublished,
			topics:       nil,
			publishedAt:  time.Now(),
			lastEditedAt: time.Now(),
		},
		{
			name:        "invalid_slug",
			expectedErr: "slug cannot be invalid",

			id:           id01,
			title:        gofakeit.Sentence(40),
			slug:         news.Slug{},
			content:      gofakeit.Sentence(500),
			status:       news.StatusPublished,
			topics:       nil,
			publishedAt:  time.Now(),
			lastEditedAt: time.Now(),
		},
		{
			name:        "invalid_status",
			expectedErr: "status cannot be invalid",

			id:           id01,
			title:        gofakeit.Sentence(40),
			slug:         news.MustNewSlugFromID(id01, slug.Make(gofakeit.Quote())),
			content:      gofakeit.Sentence(500),
			status:       news.Status{},
			topics:       nil,
			publishedAt:  time.Now(),
			lastEditedAt: time.Now(),
		},
		{
			name:        "invalid_published_at",
			expectedErr: "publishedAt cannot be invalid",

			id:           id01,
			title:        gofakeit.Sentence(40),
			slug:         news.MustNewSlugFromID(id01, slug.Make(gofakeit.Quote())),
			content:      gofakeit.Sentence(500),
			status:       news.StatusDraft,
			topics:       nil,
			publishedAt:  time.Time{},
			lastEditedAt: time.Now(),
		},
		{
			name:        "invalid_lastEditedAt",
			expectedErr: "lastEditedAt cannot be invalid",

			id:           id01,
			title:        gofakeit.Sentence(40),
			slug:         news.MustNewSlugFromID(id01, slug.Make(gofakeit.Quote())),
			content:      gofakeit.Sentence(500),
			status:       news.StatusDraft,
			topics:       nil,
			publishedAt:  time.Now(),
			lastEditedAt: time.Time{},
		},
		{
			name:        "create_news",
			expectedErr: "",

			id:           id01,
			title:        gofakeit.Sentence(40),
			slug:         news.MustNewSlugFromID(id01, slug.Make(gofakeit.Quote())),
			content:      gofakeit.Sentence(500),
			status:       news.StatusPublished,
			topics:       nil,
			publishedAt:  time.Now(),
			lastEditedAt: time.Now(),
		},
	}

	for idx := range testCases {
		tc := testCases[idx]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			n, err := news.New(
				tc.id,
				tc.title,
				tc.slug,
				tc.content,
				tc.status,
				tc.topics,
				tc.publishedAt,
				tc.lastEditedAt,
			)

			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				return
			}

			assert.Equal(t, tc.id, n.Id())
			assert.Equal(t, tc.title, n.Title())
			assert.Equal(t, tc.slug, n.Slug())
			assert.Equal(t, tc.content, n.Content())
			assert.Equal(t, tc.status, n.Status())
			assert.Equal(t, len(tc.topics), len(n.Topics()))
			assert.Equal(t, tc.publishedAt, n.PublishedAt())
			assert.Equal(t, tc.lastEditedAt, n.LastEditedAt())
		})
	}
}
