package users_service

import (
	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	usersStorage storage.UsersStorage
}

func NewUsersService(storage storage.UsersStorage) *UsersService {
	return &UsersService{usersStorage: storage}
}

func (u *UsersService) CreateUser(user *model.User) (int, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return 0, nil
	}
	user.Password = string(hash)
	return u.usersStorage.CreateUser(user)
}
