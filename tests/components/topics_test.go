package components

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/flowck/news_service_ddd_golang/tests/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTopics(t *testing.T) {
	// create topic
	res01, err := getClient(t).CreateTopic(ctx, fixtureTopic(t))
	require.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res01.StatusCode)

	// get topics
	res02, err := getClient(t).GetTopicsWithResponse(ctx)
	require.Nil(t, err)
	assert.Equal(t, http.StatusOK, res02.StatusCode())

	for _, topic := range res02.JSON200.Data {
		assert.NotEmpty(t, topic.Id)
		assert.NotEmpty(t, topic.Name)
	}
}

func fixtureTopic(t *testing.T) client.CreateTopicRequest {
	return client.CreateTopicRequest{Name: gofakeit.BeerName()}
}
