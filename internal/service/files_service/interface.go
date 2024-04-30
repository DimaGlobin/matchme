package files_service

import (
	"io"

	"github.com/DimaGlobin/matchme/internal/model"
)

type FilesServiceInt interface {
	UploadFile(fd *model.FileData, file io.Reader) (uint64, error)
	GetFileById(fileId, userId uint64) (*model.File, error)
	GetFileByName(userId uint64, filename string) (*model.File, error)
	DeleteFile(fileId, userId uint64) error
	GetAllFiles(userId uint64) ([]*model.File, error)
}

var _ FilesServiceInt = (*FilesService)(nil)
