package cache_storage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheRedis struct {
	rdb *redis.Client
}

func NewCacheRedis(rdb *redis.Client) *CacheRedis {
	return &CacheRedis{
		rdb: rdb,
	}
}

func (c *CacheRedis) AddToCache(ctx context.Context, key string, val interface{}) error {
	return c.rdb.Set(ctx, key, val, 10*time.Minute).Err()
}

func (c *CacheRedis) GetFromCache(ctx context.Context, key string) (string, error) {
	val, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}
