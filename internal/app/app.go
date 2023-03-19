package app

import (
	"context"

	"github.com/flowck/news_service_ddd_golang/internal/app/commands"
	"github.com/flowck/news_service_ddd_golang/internal/app/queries"
	"github.com/flowck/news_service_ddd_golang/internal/domain/news"
)

type CommandHandler[C any] interface {
	Execute(ctx context.Context, cmd C) error
}

type QueryHandler[Q any, R any] interface {
	Execute(ctx context.Context, q Q) (R, error)
}

type Commands struct {
	PublishNews     CommandHandler[commands.PublishNews]
	UnPublishNews   CommandHandler[commands.UnPublishNews]
	EditNews        CommandHandler[any]
	CreateTopic     CommandHandler[commands.CreateTopic]
	RemoveTopicByID CommandHandler[commands.RemoveTopicByID]
	EditTopic       CommandHandler[commands.EditTopic]
}

type Queries struct {
	NewsByID      QueryHandler[queries.NewsByID, *news.News]
	NewsByFilters QueryHandler[queries.NewsByFilters, queries.NewsPaginated]
	TopicByID     QueryHandler[queries.TopicByID, *news.Topic]
	Topics        QueryHandler[queries.Topics, []*news.Topic]
}

type App struct {
	Commands Commands
	Queries  Queries
}
