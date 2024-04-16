package files_storage

import (
	"io"

	"github.com/DimaGlobin/matchme/internal/model"
)

type FilesStorage interface {
	UploadFile(fd *model.FileData, file io.Reader) error
	GetFile(userId uint64, fd *model.FileData) ([]byte, error)
	DeleteFile(userId uint64, fd *model.FileData) error
}

var _ FilesStorage = (*FilesMinio)(nil)
