package postgres

import "github.com/jmoiron/sqlx"

type Postgres struct{}

func New(pgConn *sqlx.DB) *Postgres {
	return &Postgres{}
}
