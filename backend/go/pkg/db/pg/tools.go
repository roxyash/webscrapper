package pg

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

func setDefaultConfigValues(config *Config) {
	if config.MaxIdleConns == 0 {
		config.MaxIdleConns = 10 // default value
	}
	if config.MaxOpenConns == 0 {
		config.MaxOpenConns = 100 // default value
	}
	if config.MaxConnLifeTime == 0 {
		config.MaxConnLifeTime = time.Hour // default value
	}
	if config.MaxIdleLifeTime == 0 {
		config.MaxIdleLifeTime = 30 * time.Minute // default value
	}
	if config.MaxRetries == 0 {
		config.MaxRetries = 10 // default value
	}
	if config.RetryInterval == 0 {
		config.RetryInterval = time.Second // default value
	}
}

func setDBConnectionProperties(db *sqlx.DB, config Config) {
	db.SetConnMaxIdleTime(config.MaxIdleLifeTime)
	db.SetConnMaxLifetime(config.MaxConnLifeTime)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)
}

func getDsn(config Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)
}
