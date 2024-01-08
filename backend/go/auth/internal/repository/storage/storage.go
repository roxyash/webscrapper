package storage

import (
	"github.com/jmoiron/sqlx"
	"webscrapper/auth/internal/repository/storage/postgres"
)

type Storage struct {
	Postgres *postgres.Postgres
}

func New(pgConn *sqlx.DB) *Storage {
	return &Storage{
		Postgres: postgres.New(pgConn),
	}
}
