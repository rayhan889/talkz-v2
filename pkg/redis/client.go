package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Background = context.Background()

func InitRedisClient(addr string, password string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	err := client.Ping(Background).Err()
	if err != nil {
		return nil, err
	}

	return client, nil
}
