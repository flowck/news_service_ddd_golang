package newsarticles_test

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"github.com/flowck/news_service_ddd_golang/internal/app/queries"

	"github.com/stretchr/testify/assert"

	"github.com/gosimple/slug"

	"github.com/brianvoe/gofakeit"
	"github.com/flowck/news_service_ddd_golang/internal/adapters/newsarticles"
	"github.com/flowck/news_service_ddd_golang/internal/common/psql"
	"github.com/flowck/news_service_ddd_golang/internal/domain"
	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
	"github.com/stretchr/testify/require"

	_ "github.com/lib/pq"
)

var (
	db          *sql.DB
	ctx         context.Context
	testTimeout = time.Minute * 2
)

func TestNewsRepositoryLifecycle(t *testing.T) {
	repo, err := newsarticles.NewPsqlRepository(db)
	require.Nil(t, err)

	// Insert
	n := fixtureNews(t)
	assert.Nil(t, repo.Insert(ctx, n, nil))

	// Find
	n1, err := repo.Find(ctx, n.Id())
	require.Nil(t, err)

	assert.Equal(t, n.Id(), n1.Id())
	assert.Equal(t, n.Title(), n1.Title())
	assert.Equal(t, n.Slug(), n1.Slug())

	_, err = repo.Find(ctx, domain.NewID())
	assert.ErrorIs(t, err, news.ErrNewsNotFound)

	t.Run("find_by_filters", func(t *testing.T) {
		filter01 := fixtureNewsFilters(10, 1, news.StatusPublished.String(), "")
		paginatedNews01, err := repo.FindNewsByFiltersWithPagination(ctx, filter01)
		require.Nil(t, err)

		for i := range paginatedNews01.Data {
			assert.Equal(t, news.StatusPublished, paginatedNews01.Data[i].Status())
		}

		filter02 := fixtureNewsFilters(10, 1, "", "")
		_, err = repo.FindNewsByFiltersWithPagination(ctx, filter02)
		assert.Nil(t, err)
	})

	// Update
	n.UnPublish()
	assert.Nil(t, repo.Update(ctx, n))
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

func fixtureNews(t *testing.T) *news.News {
	ID := domain.NewID()
	sl, err := news.NewSlugFromID(ID, slug.Make(gofakeit.Quote()))
	require.Nil(t, err)

	n, err := news.New(domain.NewID(),
		gofakeit.Quote(),
		sl,
		gofakeit.Sentence(400),
		news.StatusScheduled,
		nil,
		time.Now(),
		time.Now(),
	)
	require.Nil(t, err)

	return n
}

func fixtureNewsFilters(limit, page int, status, topic string) queries.NewsFilter {
	return queries.NewsFilter{
		Limit:  limit,
		Page:   page,
		Status: status,
		Topic:  topic,
	}
}
