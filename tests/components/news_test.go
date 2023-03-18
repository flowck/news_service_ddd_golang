package components

import (
	"net/http"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
	"github.com/gosimple/slug"

	"github.com/flowck/news_service_ddd_golang/tests/client"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNews(t *testing.T) {
	// t.Run("get_news", func(t *testing.T) {
	// 	res01, err := getClient(t).GetNewsWithResponse(ctx, &client.GetNewsParams{})
	// 	require.Nil(t, err)
	//
	// 	assert.Equal(t, http.StatusOK, res01.StatusCode())
	// })

	// publish_news
	news01 := fixtureNews()
	res01, err := getClient(t).PublishNews(ctx, news01)
	require.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res01.StatusCode)
}

func fixtureNews() client.PublishNewsJSONRequestBody {
	return client.PublishNewsJSONRequestBody{
		Content:     gofakeit.Sentence(500),
		PublishedAt: time.Now(),
		Slug:        slug.Make(gofakeit.Sentence(10)),
		Status:      news.StatusPublished.String(),
		Title:       gofakeit.Sentence(20),
		TopicIds:    []string{},
	}
}
