package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"webscrapper/auth/internal/repository/cache"
	"webscrapper/auth/internal/repository/storage"
)

type Repository struct {
	Storage *storage.Storage
	Cache   *cache.Cache
}

func New(pgConn *sqlx.DB, rdConn *redis.Client) *Repository {
	return &Repository{
		Storage: storage.New(pgConn),
		Cache:   cache.New(rdConn),
	}
}
