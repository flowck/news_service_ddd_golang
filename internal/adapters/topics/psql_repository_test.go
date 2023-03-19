package topics_test

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flowck/news_service_ddd_golang/internal/adapters/topics"
	"github.com/flowck/news_service_ddd_golang/internal/common/psql"
	"github.com/flowck/news_service_ddd_golang/internal/domain"
	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
)

var (
	db          *sql.DB
	ctx         context.Context
	testTimeout = time.Minute * 2
)

func TestTopicsRepositoryLifecycle(t *testing.T) {
	repo, err := topics.NewPsqlRepository(db)
	require.Nil(t, err)

	// Insert
	topic := fixtureTopic(t)
	assert.Nil(t, repo.Insert(ctx, topic))

	// Seed
	for i := 0; i < 5; i++ {
		require.Nil(t, repo.Insert(ctx, fixtureTopic(t)))
	}

	// FindAll
	topics, err := repo.FindAll(ctx)
	require.Nil(t, err)
	assert.Equal(t, true, len(topics) > 0)

	// Find
	topic00, err := repo.Find(ctx, topics[0].ID())
	require.Nil(t, err)

	assert.Equal(t, topics[0].ID(), topic00.ID())
	assert.Equal(t, topics[0].Value(), topic00.Value())

	// Delete by ID
	err = repo.DeleteByID(ctx, topic00.ID())
	assert.Nil(t, err)

	err = repo.DeleteByID(ctx, domain.NewID())
	assert.ErrorIs(t, err, news.ErrTopicNotFound)

	// Update
	toBeUpdatedTopic := topics[1]
	require.Nil(t, toBeUpdatedTopic.Edit(gofakeit.BeerName()))

	err = repo.Update(ctx, toBeUpdatedTopic)
	assert.Nil(t, err)

	updatedTopic, err := repo.Find(ctx, toBeUpdatedTopic.ID())
	require.Nil(t, err)
	assert.Equal(t, toBeUpdatedTopic.Value(), updatedTopic.Value())
}

func TestMain(m *testing.M) {
	gofakeit.Seed(0)

	tmpctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()
	ctx = tmpctx

	var err error
	db, err = psql.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	if err = psql.ApplyMigrations(db, "../../../misc/sql/migrations"); err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func fixtureTopic(t *testing.T) *news.Topic {
	n, err := news.NewTopic(
		domain.NewID(),
		gofakeit.BeerName(),
	)
	require.Nil(t, err)

	return n
}
