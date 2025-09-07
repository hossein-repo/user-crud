package redis

import (
	"context"
	"fmt"
	"user-crud/config"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func InitRedis(cfg *config.RedisConfig) error {
	Client = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx := context.Background()
	_, err := Client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	return nil
}
