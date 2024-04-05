package storage

import (
	"context"
	"io"

	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/storage/files_data_storage"
	"github.com/DimaGlobin/matchme/internal/storage/files_storage"
	"github.com/DimaGlobin/matchme/internal/storage/ratings_storage"
	"github.com/DimaGlobin/matchme/internal/storage/users_storage"
	"github.com/jmoiron/sqlx"
)

type UsersStorage interface {
	CreateUser(user *model.User) (uint64, error)
	GetRandomUser(userId uint64) (*model.User, error)
	GetUser(email string) (*model.User, error)
	GetUserById(id uint64) (*model.User, error)
	UpdateUser(id uint64, updates model.Updates) error
	DeleteUser(id uint64) error
}

type FilesStorage interface {
	UploadFile(ctx context.Context, fd *model.FileData, file io.Reader) error
	GetFile(ctx context.Context, userId uint64, fd *model.FileData) ([]byte, error)
	DeleteFile(ctx context.Context, userId uint64, fd *model.FileData) error
}

type FilesDataStorage interface {
	AddFile(data *model.FileData) (uint64, error)
	GetFileById(fileId, userId uint64) (*model.FileData, error)
	GetFileByName(userId uint64, filename string) (*model.FileData, error)
	GetAllFiles(userId uint64) ([]*model.FileData, error)
	DeleteFile(fileId, userId uint64) error
	GetFilesCount(userId uint64) (int, error)
}

type RatingsStorage interface {
	AddLike(liking, liked uint64) (uint64, error)
	AddDislike(liking, liked uint64) (uint64, error)
	GetAllLikes(userId uint64) ([]uint64, error)
}

var _ UsersStorage = (*users_storage.UserPostgres)(nil)
var _ FilesStorage = (*files_storage.FilesMinio)(nil)
var _ FilesDataStorage = (*files_data_storage.FilesPostgres)(nil)
var _ RatingsStorage = (*ratings_storage.RatingsPostgres)(nil)

type Storage struct {
	UsersStorage
	FilesStorage
	FilesDataStorage
	RatingsStorage
}

func NewStorage(db *sqlx.DB, minioClient *files_storage.MinioClient) *Storage {
	return &Storage{
		UsersStorage:     users_storage.NewUsersPostgres(db),
		FilesStorage:     files_storage.NewFilesMinio(minioClient),
		FilesDataStorage: files_data_storage.NewFilesPostgres(db),
		RatingsStorage:   ratings_storage.NewRatingsPostgres(db),
	}
}
