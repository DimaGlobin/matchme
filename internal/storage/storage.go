package storage

import (
	"context"
	"io"

	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/storage/files_data_storage"
	"github.com/DimaGlobin/matchme/internal/storage/files_storage"
	"github.com/DimaGlobin/matchme/internal/storage/users_storage"
	"github.com/jmoiron/sqlx"
)

type UsersStorage interface {
	CreateUser(user *model.User) (int, error)
	GetUser(email string) (*model.User, error)
	GetUserById(id int) (*model.User, error)
	UpdateUser(id int, updates model.Updates) error
	DeleteUser(id int) error
}

type FilesStorage interface {
	UploadFile(ctx context.Context, fd *model.FileData, file io.Reader) error
}

type FilesDataStorage interface {
	AddFile(data *model.FileData) (int, error)
}

var _ UsersStorage = (*users_storage.UserPostgres)(nil)
var _ FilesStorage = (*files_storage.FilesMinio)(nil)
var _ FilesDataStorage = (*files_data_storage.FilesPostgres)(nil)

type Storage struct {
	UsersStorage
	FilesStorage
	FilesDataStorage
}

func NewStorage(db *sqlx.DB, minioClient *files_storage.MinioClient) *Storage {
	return &Storage{
		UsersStorage:     users_storage.NewUsersPostgres(db),
		FilesStorage:     files_storage.NewFilesMinio(minioClient),
		FilesDataStorage: files_data_storage.NewFilesPostgres(db),
	}
}
