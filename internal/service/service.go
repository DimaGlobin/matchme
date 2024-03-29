package service

import (
	"context"
	"io"

	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/service/files_service"
	"github.com/DimaGlobin/matchme/internal/service/users_service"
	"github.com/DimaGlobin/matchme/internal/storage"
)

type UsersService interface {
	CreateUser(user *model.User) (uint64, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (uint64, error)
	GetuserById(id uint64) (*model.User, error)
	UpdateUser(id uint64, updates model.Updates) error
	DeleteUser(id uint64) error
}

type FilesService interface {
	UploadFile(ctx context.Context, fd *model.FileData, file io.Reader) (uint64, error)
	GetFile(ctx context.Context, fileId, userId uint64) (*model.File, error)
}

var _ UsersService = (*users_service.UsersService)(nil)
var _FilesService = (*files_service.FilesService)(nil)

type Service struct {
	UsersService
	FilesService
}

func NewService(storage storage.Storage) *Service {
	return &Service{
		UsersService: users_service.NewUsersService(storage.UsersStorage),
		FilesService: files_service.NewFilesService(storage.FilesStorage, storage.FilesDataStorage),
	}
}
