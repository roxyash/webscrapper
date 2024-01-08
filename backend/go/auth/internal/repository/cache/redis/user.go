package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
	"webscrapper/auth/internal/model"
)

type User interface {
	GetByUsername(username string) (model.User, error)
	Set(user model.User) error
}

type user struct {
	conn       *redis.Client
	tag        string
	defaultTTL time.Duration
}

func NewUser(conn *redis.Client) User {
	return &user{
		conn:       conn,
		tag:        "user",
		defaultTTL: time.Hour * 1, // 1 hour
	}
}

func (r *user) GetByUsername(username string) (model.User, error) {
	// Get user from redis
	val, err := r.conn.Get(context.Background(), r.tag+":"+username).Result()
	if err != nil {
		return model.User{}, err
	}

	// Unmarshal user from json
	var user model.User
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *user) Set(user model.User) error {
	// Marshal user to json
	userJson, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// Set user to redis
	err = r.conn.Set(context.Background(), r.tag+":"+user.Username, userJson, r.defaultTTL).Err()
	if err != nil {
		return err
	}

	return nil
}
