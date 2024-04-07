package service

import (
	"context"
	"io"

	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/service/files_service"
	"github.com/DimaGlobin/matchme/internal/service/ratings_service"
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
	GetFileById(ctx context.Context, fileId, userId uint64) (*model.File, error)
	GetFileByName(ctx context.Context, userId uint64, filename string) (*model.File, error)
	DeleteFile(ctx context.Context, fileId, userId uint64) error
	GetAllFiles(ctx context.Context, userId uint64) ([]*model.File, error)
}

type RatingsService interface {
	RecommendUser(userId uint64) (*model.User, error)
	AddReaction(reaction string, subjectId, objectId uint64) (uint64, uint64, error)
	GetAllLikes(userId uint64) ([]uint64, error)
}

var _ UsersService = (*users_service.UsersService)(nil)
var _ FilesService = (*files_service.FilesService)(nil)
var _ RatingsService = (*ratings_service.RatingsService)(nil)

type Service struct {
	UsersService
	FilesService
	RatingsService
}

func NewService(storage storage.Storage) *Service {
	return &Service{
		UsersService:   users_service.NewUsersService(storage.UsersStorage),
		FilesService:   files_service.NewFilesService(storage.FilesStorage, storage.FilesDataStorage),
		RatingsService: ratings_service.NewRatingsService(storage.RatingsStorage, storage.UsersStorage),
	}
}
