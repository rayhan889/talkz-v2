package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var Background = context.Background()

func InitRedisClient(addr string, password string, db int) error {
	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	err := Client.Ping(Background).Err()
	if err != nil {
		return err
	}

	return nil

}
