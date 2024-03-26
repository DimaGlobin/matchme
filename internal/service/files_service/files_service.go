package files_service

import (
	"context"
	"io"

	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/storage"
)

type FilesService struct {
	filesStorage storage.FilesStorage
}

func NewFilesService(filesStorage storage.FilesStorage) *FilesService {
	return &FilesService{
		filesStorage: filesStorage,
	}
}

func (f *FilesService) UploadFile(ctx context.Context, fd *model.FileData, file io.Reader) error {
	return f.filesStorage.UploadFile(ctx, fd, file)
}
