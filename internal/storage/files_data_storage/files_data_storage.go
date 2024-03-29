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

func (f *FilesPostgres) AddFile(data *model.FileData) (int, error) {
	var id int

	query := "INSERT INTO files (user_id, filename, size) values ($1, $2, $3) RETURNING file_id"
	row := f.db.QueryRow(query, data.UserId, data.FileName, data.Size)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
