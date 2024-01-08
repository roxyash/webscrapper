package repository

import (
	"github.com/jmoiron/sqlx"
	"webscrapper/api-gateway/internal/repository/cache"
	"webscrapper/api-gateway/internal/repository/storage"
	"webscrapper/api-gateway/internal/repository/webapi"
)

type Repository struct {
	WebApi  *webapi.WebApi
	Storage *storage.Storage
	Cache   *cache.Cache
}

func New(pgConn *sqlx.DB) *Repository {
	return &Repository{
		WebApi:  webapi.New(),
		Storage: storage.New(pgConn),
		Cache:   cache.New(),
	}
}
