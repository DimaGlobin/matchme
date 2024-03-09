package storage

import (
	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/storage/users_storage"
	"github.com/jmoiron/sqlx"
)

type UsersStorage interface {
	CreateUser(user *model.User) (int, error)
	GetUser(email string) (*model.User, error)
}

type PhotoStorage interface {
	AddPhoto() error
	DeletePhoto()
}

var _ UsersStorage = (*users_storage.UserPostgres)(nil)

type Storage struct {
	UsersStorage
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		UsersStorage: users_storage.NewUsersPostgres(db),
	}
}
