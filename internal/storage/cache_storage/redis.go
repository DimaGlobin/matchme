package cache_storage

import (
	"context"
	"fmt"

	"github.com/DimaGlobin/matchme/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisDB(cfg *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.CacheStorage.Host, cfg.CacheStorage.Port),
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
