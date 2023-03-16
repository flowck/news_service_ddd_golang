package news

import "context"

type Repository interface {
	Insert(ctx context.Context, n News) error
}
