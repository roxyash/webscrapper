package redis

import "github.com/go-redis/redis/v8"

type Redis struct {
	User
}

func New(conn *redis.Client) *Redis {
	return &Redis{
		User: NewUser(conn),
	}
}
