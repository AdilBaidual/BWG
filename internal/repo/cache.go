package repo

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) *RedisRepository {
	return &RedisRepository{
		client: client,
	}
}

func (r *RedisRepository) Set(key string, value string) error {
	ctx := context.Background()
	return r.client.Set(ctx, key, value, time.Minute).Err()
}

func (r *RedisRepository) Get(key string) (string, error) {
	ctx := context.Background()
	return r.client.Get(ctx, key).Result()
}
