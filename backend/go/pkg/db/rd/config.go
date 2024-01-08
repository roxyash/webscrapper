package rd

import "time"

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DB       int

	MaxRetries    int           // optional default 10
	RetryInterval time.Duration // optional default 1 second
}

func (c *Config) GetAddr() string {
	return c.Host + ":" + c.Port
}
