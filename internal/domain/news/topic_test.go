package news_test

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/flowck/news_service_ddd_golang/internal/domain"
	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTopic(t *testing.T) {
	testCases := []struct {
		name        string
		expectedErr string

		ID    domain.ID
		value string
	}{
		{
			name:        "create_topic",
			expectedErr: "",

			ID:    domain.NewID(),
			value: gofakeit.BuzzWord(),
		},
		{
			name:        "invalid_id",
			expectedErr: "id cannot be invalid",

			ID:    domain.ID{},
			value: gofakeit.BuzzWord(),
		},
		{
			name:        "create_topic",
			expectedErr: "value cannot be empty",

			ID:    domain.NewID(),
			value: "",
		},
	}

	for idx := range testCases {
		tc := testCases[idx]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			topic, err := news.NewTopic(tc.ID, tc.value)

			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				return
			}

			require.Nil(t, err)
			assert.Equal(t, tc.ID, topic.ID())
			assert.Equal(t, tc.value, topic.Value())
		})
	}
}
