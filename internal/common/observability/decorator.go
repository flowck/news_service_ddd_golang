package observability

import (
	"context"
	"fmt"
	"strings"
)

type QueryHandler[Q any, R any] interface {
	Execute(ctx context.Context, q Q) (result R, err error)
}

type CommandHandler[C any] interface {
	Execute(ctx context.Context, cmd C) error
}

func handlerName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}

func prettyPrint(v any) string {
	return fmt.Sprintf("%v", v)
}
