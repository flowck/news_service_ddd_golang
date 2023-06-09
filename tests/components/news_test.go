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
	// publish_news
	news01 := fixtureNews()
	res01, err := getClient(t).PublishNews(ctx, news01)
	require.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res01.StatusCode)

	// seed_news
	for i := 0; i < 10; i++ {
		_, err = getClient(t).PublishNews(ctx, fixtureNews())
		require.Nil(t, err)
	}

	// get_news_by_filters_and_pagination
	limit01 := 5
	res02, err := getClient(t).GetNewsWithResponse(ctx, &client.GetNewsParams{
		Limit: toPtr(limit01),
		Page:  toPtr(1),
	})

	require.Nil(t, err)
	assert.Equal(t, http.StatusOK, res02.StatusCode())
	assert.Equal(t, limit01, len(res02.JSON200.Data))
	assert.Equal(t, 1, res02.JSON200.Page)

	newsListResponse := res02.JSON200.Data

	// get_news_by_id
	news01byId, err := getClient(t).GetNewsByIDWithResponse(ctx, res02.JSON200.Data[0].Id)

	require.Nil(t, err)
	assert.Equal(t, http.StatusOK, news01byId.StatusCode())
	assert.Equal(t, res02.JSON200.Data[0].Id, news01byId.JSON200.Data.Id)
	assert.Equal(t, res02.JSON200.Data[0].Title, news01byId.JSON200.Data.Title)
	assert.Equal(t, res02.JSON200.Data[0].Status, news01byId.JSON200.Data.Status)
	assert.Equal(t, res02.JSON200.Data[0].Slug, news01byId.JSON200.Data.Slug)

	// un_publish_news
	res03, err := getClient(t).UnPublishNews(ctx, res02.JSON200.Data[0].Id)
	require.Nil(t, err)
	assert.Equal(t, http.StatusNoContent, res03.StatusCode)

	unPublishedNews, err := getClient(t).GetNewsByIDWithResponse(ctx, res02.JSON200.Data[0].Id)
	require.Nil(t, err)
	assert.Equal(t, news.StatusDraft.String(), unPublishedNews.JSON200.Data.Status)

	// edit_news
	editNewsReq := client.EditNewsJSONRequestBody{
		Content:     gofakeit.Sentence(30),
		PublishedAt: time.Now().Add(time.Hour * 24),
		Slug:        slug.Make(gofakeit.Question()),
		Status:      news.StatusPublished.String(),
		Title:       gofakeit.Question(),
		TopicsIds:   nil,
	}
	res04, err := getClient(t).EditNews(ctx, newsListResponse[2].Id, editNewsReq)

	require.Nil(t, err)
	assert.Equal(t, http.StatusNoContent, res04.StatusCode)

	// get_news_by_id
	editedNews, err := getClient(t).GetNewsByIDWithResponse(ctx, newsListResponse[2].Id)

	require.Nil(t, err)
	assert.Equal(t, http.StatusOK, editedNews.StatusCode())
	assert.Equal(t, editNewsReq.Title, editedNews.JSON200.Data.Title)
	assert.Equal(t, editNewsReq.Status, editedNews.JSON200.Data.Status)
	assert.Equal(t, editNewsReq.Slug, editedNews.JSON200.Data.Slug)
	assert.Equal(t, editNewsReq.Content, editedNews.JSON200.Data.Content)
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
