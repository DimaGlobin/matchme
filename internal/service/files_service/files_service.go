package files_service

import (
	"context"
	"io"

	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/storage"
)

type FilesService struct {
	filesStorage     storage.FilesStorage
	filesDataStorage storage.FilesDataStorage
}

func NewFilesService(filesStorage storage.FilesStorage, filesDataStorage storage.FilesDataStorage) *FilesService {
	return &FilesService{
		filesStorage:     filesStorage,
		filesDataStorage: filesDataStorage,
	}
}

func (f *FilesService) UploadFile(ctx context.Context, fd *model.FileData, file io.Reader) (int, error) {
	if err := f.filesStorage.UploadFile(ctx, fd, file); err != nil {
		return 0, err
	}

	return f.filesDataStorage.AddFile(fd)
}
