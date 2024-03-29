package files_data_storage

import (
	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/jmoiron/sqlx"
)

type FilesPostgres struct {
	db *sqlx.DB
}

func NewFilesPostgres(db *sqlx.DB) *FilesPostgres {
	return &FilesPostgres{db: db}
}

func (f *FilesPostgres) AddFile(data *model.FileData) (uint64, error) {
	var id uint64

	query := "INSERT INTO files (user_id, file_name, size) values ($1, $2, $3) RETURNING file_id"
	row := f.db.QueryRow(query, data.UserId, data.FileName, data.Size)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (f *FilesPostgres) GetFile(fileId, userId uint64) (*model.FileData, error) {
	fd := &model.FileData{}

	query := "SELECT * FROM files WHERE file_id=$1 AND user_id=$2"

	if err := f.db.Get(fd, query, fileId, userId); err != nil {
		return nil, err
	}

	return fd, nil
}
