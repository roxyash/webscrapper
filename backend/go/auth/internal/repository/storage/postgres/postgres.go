package postgres

import (
	"github.com/jmoiron/sqlx"
	"webscrapper/auth/internal/config"
	"webscrapper/pkg/logging"
)

type Postgres struct {
	User
}

var (
	conf   = config.GetConfig()
	logger = logging.GetLogger()
)

func New(pgConn *sqlx.DB) *Postgres {
	schema := conf.Postgres.Schema

	return &Postgres{
		User: NewUser(pgConn, schema),
	}
}
