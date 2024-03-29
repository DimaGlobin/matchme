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

func (f *FilesService) UploadFile(ctx context.Context, fd *model.FileData, file io.Reader) (uint64, error) {
	if err := f.filesStorage.UploadFile(ctx, fd, file); err != nil {
		return 0, err
	}

	return f.filesDataStorage.AddFile(fd)
}

func (f *FilesService) GetFile(ctx context.Context, fileId, userId uint64) (*model.File, error) {
	fd, err := f.filesDataStorage.GetFile(fileId, userId)
	if err != nil {
		return nil, err
	}

	buf, err := f.filesStorage.GetFile(ctx, userId, fd)
	if err != nil {
		return nil, err
	}

	return &model.File{
		Name:   fd.FileName,
		Buffer: buf,
	}, nil
}
