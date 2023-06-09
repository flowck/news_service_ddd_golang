package main

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"

	"github.com/flowck/news_service_ddd_golang/internal/adapters/newsarticles"
	"github.com/flowck/news_service_ddd_golang/internal/adapters/topics"
	"github.com/flowck/news_service_ddd_golang/internal/adapters/transaction"
	"github.com/flowck/news_service_ddd_golang/internal/app"
	"github.com/flowck/news_service_ddd_golang/internal/app/commands"
	"github.com/flowck/news_service_ddd_golang/internal/app/queries"
	"github.com/flowck/news_service_ddd_golang/internal/common/logs"
	"github.com/flowck/news_service_ddd_golang/internal/common/observability"
	"github.com/flowck/news_service_ddd_golang/internal/common/psql"
	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
	"github.com/flowck/news_service_ddd_golang/internal/ports/http"
)

type Config struct {
	Port              int16  `envconfig:"PORT"`
	DebugMode         string `envconfig:"FLAG_DEBUG_MODE"`
	AllowedCorsOrigin string `envconfig:"ALLOWED_CORS_ORIGIN"`
	DatabaseUrl       string `envconfig:"DATABASE_URL"`
}

func main() {
	cfg, err := getConfig()
	if err != nil {
		panic(err)
	}

	logger := logs.New(cfg.DebugMode == "enabled")
	logger.Info("news_service")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	db, err := psql.Connect(cfg.DatabaseUrl)
	if err != nil {
		logger.Fatal(err)
	}

	err = psql.ApplyMigrations(db, "misc/sql/migrations")
	if err != nil {
		logger.Fatal(err)
	}

	newsRepository, err := newsarticles.NewPsqlRepository(db)
	if err != nil {
		logger.Fatal(err)
	}

	topicsRepository, err := topics.NewPsqlRepository(db)
	if err != nil {
		logger.Fatal(err)
	}

	txProvider := transaction.NewSQLProvider(db)

	application := &app.App{
		Commands: app.Commands{
			PublishNews:     observability.NewCommandDecorator[commands.PublishNews](commands.NewPublishNewsHandler(txProvider), logger),
			UnPublishNews:   commands.NewUnPublishNewsHandler(txProvider),
			EditNews:        commands.NewEditNewsHandler(newsRepository),
			CreateTopic:     commands.NewCreateTopicHandler(topicsRepository),
			RemoveTopicByID: commands.NewRemoveTopicByIDHandler(topicsRepository),
			EditTopic:       commands.NewEditTopicHandler(topicsRepository),
		},
		Queries: app.Queries{
			NewsByID:      observability.NewQueryDecorator[queries.NewsByID, *news.News](queries.NewNewsByIdHandler(newsRepository), logger),
			NewsByFilters: queries.NewNewsByFiltersHandler(newsRepository),
			TopicByID:     queries.NewTopicByIdHandler(topicsRepository),
			Topics:        queries.NewTopicsHandler(topicsRepository),
		},
	}

	httpPort := http.NewPort(ctx, cfg.Port, strings.Split(cfg.AllowedCorsOrigin, ";"), application, logger)
	httpPort.Start()

	<-done
	logger.Info("Preparing to shutdown gracefully")

	ctxTerm, cancelCtxTerm := context.WithTimeout(ctx, time.Second*15)
	defer cancelCtxTerm()

	httpPort.Stop(ctxTerm)

	logger.Info("The service has been terminated")
}

func getConfig() (*Config, error) {
	cfg := &Config{}

	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
