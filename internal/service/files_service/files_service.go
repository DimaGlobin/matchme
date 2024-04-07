package files_service

import (
	"context"
	"fmt"
	"io"

	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/storage"
)

const (
	filesLimit = 6
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

	count, err := f.filesDataStorage.GetFilesCount(fd.UserId)
	if err != nil {
		return 0, err
	}
	if count >= filesLimit {
		return 0, fmt.Errorf("Files limit exceeded")
	}

	if err := f.filesStorage.UploadFile(ctx, fd, file); err != nil {
		return 0, err
	}

	return f.filesDataStorage.AddFile(fd)
}

func (f *FilesService) GetFileById(ctx context.Context, fileId, userId uint64) (*model.File, error) {
	fd, err := f.filesDataStorage.GetFileById(fileId, userId)
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

func (f *FilesService) GetFileByName(ctx context.Context, userId uint64, filename string) (*model.File, error) {
	fd, err := f.filesDataStorage.GetFileByName(userId, filename)
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

func (f *FilesService) GetAllFiles(ctx context.Context, userId uint64) ([]*model.File, error) {
	files := []*model.File{}
	filesData, err := f.filesDataStorage.GetAllFiles(userId)
	if err != nil {
		return nil, err
	}

	for _, v := range filesData {
		buf, err := f.filesStorage.GetFile(ctx, userId, v)
		if err != nil {
			return nil, err
		}
		files = append(files, &model.File{
			Name:   v.FileName,
			Buffer: buf,
		})
	}

	return files, nil
}

func (f *FilesService) DeleteFile(ctx context.Context, fileId, userId uint64) error {
	fd, err := f.filesDataStorage.GetFileById(fileId, userId)
	if err != nil {
		return err
	}

	err = f.filesStorage.DeleteFile(ctx, userId, fd)
	if err != nil {
		return err
	}

	return f.filesDataStorage.DeleteFile(fileId, userId)
}
