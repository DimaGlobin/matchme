package users_service

import (
	"github.com/DimaGlobin/matchme/internal/model"
)

type UsersServiceInt interface {
	CreateUser(user *model.User) (uint64, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(accessToken string) (*tokenClaims, error)
	GetuserById(id uint64) (*model.User, error)
	UpdateUser(id uint64, updates model.Updates) error
	DeleteUser(id uint64) error
}

var _ UsersServiceInt = (*UsersService)(nil)
