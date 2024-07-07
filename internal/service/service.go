package service

import (
	"github.com/DimaGlobin/matchme/internal/service/files_service"
	"github.com/DimaGlobin/matchme/internal/service/ratings_service"
	"github.com/DimaGlobin/matchme/internal/service/users_service"
	"github.com/DimaGlobin/matchme/internal/storage"
)

type Service struct {
	users_service.UsersServiceInt
	files_service.FilesServiceInt
	ratings_service.RatingsServiceInt
}

func NewService(storage storage.Storage) *Service {
	return &Service{
		UsersServiceInt:   users_service.NewUsersService(storage.UsersStorage, storage.CacheStorage),
		FilesServiceInt:   files_service.NewFilesService(storage.FilesStorage, storage.FilesDataStorage),
		RatingsServiceInt: ratings_service.NewRatingsService(storage.RatingsStorage, storage.UsersStorage, storage.CacheStorage),
	}
}
