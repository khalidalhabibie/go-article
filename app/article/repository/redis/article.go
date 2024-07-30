package redis

import (
	"backend/app/article"

	"github.com/go-redis/redis"
)

type Cache struct {
	cache *redis.Client
}

func New(cache *redis.Client) article.Cache {
	return &Cache{
		cache: cache,
	}
}
