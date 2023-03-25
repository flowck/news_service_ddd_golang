package observability

import (
	"context"

	"github.com/flowck/news_service_ddd_golang/internal/common/logs"
)

type commandDecorator[C any] struct {
	logger *logs.Logger
	base   CommandHandler[C]
}

func NewCommandDecorator[C any](base CommandHandler[C], logger *logs.Logger) commandDecorator[C] {
	return commandDecorator[C]{
		logger: logger,
		base: commandMetricsDecorator[C]{
			base: commandLoggingDecorator[C]{
				base:   base,
				logger: logger,
			},
		},
	}
}

func (q commandDecorator[C]) Execute(ctx context.Context, cmd C) error {
	return q.base.Execute(ctx, cmd)
}
