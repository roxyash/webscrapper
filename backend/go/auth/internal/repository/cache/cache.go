package cache

import (
	redisclient "github.com/go-redis/redis/v8"
	"webscrapper/auth/internal/repository/cache/redis"
)

type Cache struct {
	Redis *redis.Redis
}

func New(rdConn *redisclient.Client) *Cache {
	return &Cache{
		Redis: redis.New(rdConn),
	}
}
