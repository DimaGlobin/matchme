package service

import (
	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/service/users_service"
	"github.com/DimaGlobin/matchme/internal/storage"
)

type UsersService interface {
	CreateUser(user *model.User) (int, error)
}

var _ UsersService = (*users_service.UsersService)(nil)

type Service struct {
	UsersService
}

func NewService(storage storage.Storage) *Service {
	return &Service{
		UsersService: users_service.NewUsersService(storage.UsersStorage),
	}
}
