package app

import "context"

type CommandHandler[C any] interface {
	Execute(ctx context.Context, cmd C) error
}

type QueryHandler[Q any, R any] interface {
	Execute(ctx context.Context, q Q) (R, error)
}

type Commands struct {
	PublishNews   CommandHandler[any]
	UnPublishNews CommandHandler[any]
	EditNews      CommandHandler[any]
	CreateTopic   CommandHandler[any]
	RemoveTopic   CommandHandler[any]
	EditTopic     CommandHandler[any]
}

type Queries struct {
	News         QueryHandler[any, any]
	MultipleNews QueryHandler[any, any]
	Topic        QueryHandler[any, any]
	Topics       QueryHandler[any, any]
}

type App struct {
	Commands Commands
	Queries  Queries
}
