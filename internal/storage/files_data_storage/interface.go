package files_data_storage

import "github.com/DimaGlobin/matchme/internal/model"

type FilesDataStorage interface {
	AddFile(data *model.FileData) (uint64, error)
	GetFileById(fileId, userId uint64) (*model.FileData, error)
	GetFileByName(userId uint64, filename string) (*model.FileData, error)
	GetAllFiles(userId uint64) ([]*model.FileData, error)
	DeleteFile(fileId, userId uint64) error
	GetFilesCount(userId uint64) (int, error)
}

var _ FilesDataStorage = (*FilesPostgres)(nil)
