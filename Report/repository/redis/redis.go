package redis

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type reportRedis struct {
	Redis *redis.Client
}

func NewReportRedis(redis *redis.Client) *reportRedis {
	return &reportRedis{Redis: redis}
}

func (r *reportRedis) Get(ctx context.Context, key string) (interface{}, error) {
	value, err := r.Redis.Get(ctx, key).Result()

	if err != nil {
		log.Printf("redis get error")
	}

	return value, err

}

func (r *reportRedis) Set(ctx context.Context, value interface{}, key string, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		log.Printf("redis : marshaling error ")
	}
	err = r.Redis.Set(ctx, key, data, ttl).Err()

	if err != nil {
		log.Printf("redis set error")
	}

	return err

}

func (r *reportRedis) Del(ctx context.Context, key string) error {
	err := r.Redis.Del(ctx, key).Err()

	return err
}
