package components

import (
	"net/http"
	"testing"

	"github.com/flowck/news_service_ddd_golang/tests/client"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNews(t *testing.T) {
	t.Parallel()

	t.Run("get_news", func(t *testing.T) {
		res01, err := getClient(t).GetNewsWithResponse(ctx, &client.GetNewsParams{})
		require.Nil(t, err)

		assert.Equal(t, http.StatusOK, res01.StatusCode())
	})
}
