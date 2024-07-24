package cache_storage

import (
	"github.com/DimaGlobin/matchme/internal/model"
)

type CacheError string

func (c CacheError) Error() string {
	return string(c)
}

// почему здесь не завести отдельный файл с ошибками ?
const (
	NoValueInCache   CacheError = "There is no value with this key in cache"
	CannotAddToCache CacheError = "Can't add value to cache"
	DecrementError   CacheError = "Can't decrement likes count"
)

type CacheStorage interface {
	AddUserToCache(user *model.User) error
	GetUserFromCache(userId uint64) (*model.User, error)
	DecLikesCount(userId uint64) (int, error)
}

var _ CacheStorage = (*CacheRedis)(nil)
