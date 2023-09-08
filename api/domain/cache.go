package domain

import "context"

type Cache interface {
	Get(ctx context.Context, key string) (value string, err error)
}
