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

	// get topic by id
	res03, err := getClient(t).GetTopicByIDWithResponse(ctx, res02.JSON200.Data[0].Id)
	require.Nil(t, err)

	assert.Equal(t, http.StatusOK, res03.StatusCode())
	assert.Equal(t, res02.JSON200.Data[0].Id, res03.JSON200.Data.Id)
	assert.Equal(t, res02.JSON200.Data[0].Name, res03.JSON200.Data.Name)

	// topic not found
	res04, err := getClient(t).GetTopicByIDWithResponse(ctx, gofakeit.UUID())
	require.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, res04.StatusCode())

	// remove topic by id
	res05, err := getClient(t).RemoveTopicByID(ctx, res03.JSON200.Data.Id)
	require.Nil(t, err)

	assert.Equal(t, http.StatusNoContent, res05.StatusCode)

	res06, err := getClient(t).GetTopicByIDWithResponse(ctx, res03.JSON200.Data.Id)
	require.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, res06.StatusCode())
}

func fixtureTopic(t *testing.T) client.CreateTopicRequest {
	return client.CreateTopicRequest{Name: gofakeit.BeerName()}
}
