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

func (c queryLoggingDecorator[Q, R]) Execute(ctx context.Context, q Q) (result R, err error) {
	logger := c.logger.WithFields(logs.Fields{
		"name":  handlerName(c.base),
		"query": prettyPrint(q),
	})

	logger.Debug("Executing query")
	defer func() {
		if err == nil {
			logger.Info("Query executed successfully")
		} else {
			logger.WithError(err).Error("Failed to execute query")
		}
	}()

	return c.base.Execute(ctx, q)
}
