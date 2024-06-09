package cache_storage

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/redis/go-redis/v9"
)

const maxLikesPerDay = 50

type CacheRedis struct {
	rdb *redis.Client
}

func NewCacheRedis(rdb *redis.Client) *CacheRedis {
	return &CacheRedis{
		rdb: rdb,
	}
}

func (c *CacheRedis) AddUserToCache(user *model.User) error {
	ctx := context.TODO()
	key := fmt.Sprintf("user:%d", user.Id)

	userData, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return c.rdb.Set(ctx, key, userData, 10*time.Minute).Err()
}

func (c *CacheRedis) GetUserFromCache(userId uint64) (*model.User, error) {
	key := fmt.Sprintf("user:%d", userId)
	ctx := context.TODO()

	userStr, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, NoValueInCache
		}

		return nil, err
	}

	var user = &model.User{}
	err = json.Unmarshal([]byte(userStr), user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c *CacheRedis) DecLikesCount(userId uint64) (int, error) {
	key := fmt.Sprintf("users_likes:%d", userId)
	ctx := context.TODO()

	likesLeft, err := c.rdb.Get(ctx, key).Int()
	if err == redis.Nil {
		likesLeft = maxLikesPerDay
		err = c.rdb.Set(ctx, key, likesLeft, time.Until(endOfDay())).Err()
		if err != nil {
			return 0, CannotAddToCache
		}
	} else if err != nil {
		return 0, err
	}

	if likesLeft > 0 {
		err = c.rdb.Decr(ctx, key).Err()
		if err != nil {
			return 0, DecrementError
		}
		return likesLeft - 1, nil
	}

	return 0, nil
}

func endOfDay() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
}
