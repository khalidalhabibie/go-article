package redis

import (
	"backend/app/user"

	"github.com/go-redis/redis"
)

type Cache struct {
	cache *redis.Client
}

func New(cache *redis.Client) user.Cache {
	return &Cache{
		cache: cache,
	}
}
