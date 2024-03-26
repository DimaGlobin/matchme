package storage

import (
	"context"
	"io"

	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/storage/files_storage"
	"github.com/DimaGlobin/matchme/internal/storage/users_storage"
	"github.com/jmoiron/sqlx"
)

type UsersStorage interface {
	CreateUser(user *model.User) (int, error)
	GetUser(email string) (*model.User, error)
}

type FilesStorage interface {
	UploadFile(ctx context.Context, fd *model.FileData, file io.Reader) error
}

var _ UsersStorage = (*users_storage.UserPostgres)(nil)
var _ FilesStorage = (*files_storage.FilesMinio)(nil)

type Storage struct {
	UsersStorage
	FilesStorage
}

func NewStorage(db *sqlx.DB, minioClient *files_storage.MinioClient) *Storage {
	return &Storage{
		UsersStorage: users_storage.NewUsersPostgres(db),
		FilesStorage: files_storage.NewFilesMinio(minioClient),
	}
}
