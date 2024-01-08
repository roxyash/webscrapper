package pg

import (
	"fmt"
	"time"
	"webscrapper/pkg/logging"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

var (
	logger = logging.GetLogger()
)

func Connect(config Config) (*sqlx.DB, error) {
	logger.Info("try connect to postgres...")
	dsn := getDsn(config)
	fmt.Println(dsn)

	setDefaultConfigValues(&config)

	var db *sqlx.DB
	var err error

	for i := 0; i < config.MaxRetries; i++ {
		db, err = sqlx.Open("pgx", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}
		logger.Errorf("failed connect to postgres - %s, retrying - [%d] ", err.Error(), i)
		time.Sleep(config.RetryInterval)
	}

	if err != nil {
		return nil, fmt.Errorf("failed connect to postgres after [%d] retries - %w", config.MaxRetries, err)
	}

	setDBConnectionProperties(db, config)

	logger.Info("connected to postgres")
	return db, nil
}

func Close(db *sqlx.DB) {
	if err := db.Close(); err != nil {
		logger.Error("can't close postgres connection")
	}
}
