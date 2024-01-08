package cache

import "webscrapper/api-gateway/internal/repository/cache/redis"

type Cache struct {
	Redis *redis.Redis
}

func New() *Cache {
	return &Cache{
		Redis: redis.New(),
	}
}
