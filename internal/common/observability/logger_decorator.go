package observability

import (
	"context"

	"github.com/flowck/news_service_ddd_golang/internal/common/logs"
)

type commandLoggingDecorator[C any] struct {
	logger *logs.Logger
	base   CommandHandler[C]
}

func (c commandLoggingDecorator[C]) Execute(ctx context.Context, cmd C) error {
	return c.base.Execute(ctx, cmd)
}

type queryLoggingDecorator[Q any, R any] struct {
	logger *logs.Logger
	base   QueryHandler[Q, R]
}

func (c queryLoggingDecorator[Q, R]) Execute(ctx context.Context, q Q) (R, error) {
	return c.base.Execute(ctx, q)
}
