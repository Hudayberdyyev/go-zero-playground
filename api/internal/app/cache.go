package app

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Cache struct {
	client *redis.Client
}

func NewCache(client *redis.Client) *Cache {
	return &Cache{client: client}
}

func (c *Cache) Get(ctx context.Context, key string) (value string, err error) {
	return "", nil
}
