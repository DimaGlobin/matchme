package cache_storage

import (
	"github.com/DimaGlobin/matchme/internal/model"
)

type CacheError string

func (c CacheError) Error() string {
	return string(c)
}

const NoValueInCache CacheError = "There is no value with this in cache"

type CacheStorage interface {
	AddUserToCache(user *model.User) error
	GetUserFromCache(userId uint64) (*model.User, error)
}

var _ CacheStorage = (*CacheRedis)(nil)
