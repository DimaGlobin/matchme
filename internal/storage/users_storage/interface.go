package users_storage

import "github.com/DimaGlobin/matchme/internal/model"

type UsersStorage interface {
	CreateUser(user *model.User) (uint64, error)
	GetRandomUser(userId uint64) (*model.User, error)
	GetUser(email string) (*model.User, error)
	GetUserById(id uint64) (*model.User, error)
	UpdateUser(id uint64, updates model.Updates) error
	DeleteUser(id uint64) error
}

var _ UsersStorage = (*UserPostgres)(nil)
