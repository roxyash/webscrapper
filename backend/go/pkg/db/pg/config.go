package pg

import "time"

// Config is a configuration for postgres connection
// required fields: Host, Port, User, Password, Database
// optional fields: MaxIdleConns, MaxOpenConns, MaxConnLifeTime, MaxIdleLifeTime
type Config struct {
	Host     string // required
	Port     string // required
	User     string // required
	Password string // required
	Database string // required

	MaxIdleConns    int           // optional default 10
	MaxOpenConns    int           // optional default 10
	MaxConnLifeTime time.Duration // optional default 1 hour
	MaxIdleLifeTime time.Duration // optional default 30 minutes

	MaxRetries    int           // optional default 10
	RetryInterval time.Duration // optional default 1 second
}
