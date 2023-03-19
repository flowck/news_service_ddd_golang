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

	// Find
	topics, err := repo.Find(ctx)
	require.Nil(t, err)
	assert.Equal(t, true, len(topics) > 0)
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
