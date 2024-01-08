package rd

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
	"webscrapper/pkg/logging"
)

type Client struct {
	rdb *redis.Client
}

var (
	logger = logging.GetLogger()
)

func Connect(config Config) (*redis.Client, error) {
	setDefaultConfigValues(&config)

	logger.Info("try connect to redis...")
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.GetAddr(),
		Username: config.Username,
		Password: config.Password,
		DB:       config.DB,
	})

	fmt.Printf("%+v", config)

	var err error
	// Retry connection
	for i := 0; i < config.MaxRetries; i++ {
		err = rdb.Ping(context.Background()).Err()
		if err == nil {
			break
		}

		logger.Errorf("failed connect to redis: %s, retrying... ", err)
		time.Sleep(config.RetryInterval)
	}

	if err != nil {
		return nil, fmt.Errorf("failed connect to redis after %d retries: %w", config.MaxRetries, err)
	}

	logger.Info("connected to redis")
	return rdb, nil
}

func Close(client *redis.Client) {
	if err := client.Close(); err != nil {
		logger.Fatalf("failed to close redis connection: %s", err)
	}
}
