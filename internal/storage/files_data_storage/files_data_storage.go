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

func (f *FilesPostgres) GetFileById(fileId, userId uint64) (*model.FileData, error) {
	fd := &model.FileData{}

	query := "SELECT * FROM files WHERE file_id=$1 AND user_id=$2"

	if err := f.db.Get(fd, query, fileId, userId); err != nil {
		return nil, err
	}

	return fd, nil
}

func (f *FilesPostgres) GetFileByName(userId uint64, filename string) (*model.FileData, error) {
	fd := &model.FileData{}

	query := "SELECT * FROM files WHERE file_name=$1 AND user_id=$2"

	if err := f.db.Get(fd, query, filename, userId); err != nil {
		return nil, err
	}

	return fd, nil
}

func (f *FilesPostgres) GetAllFiles(userId uint64) ([]*model.FileData, error) {
	files := []*model.FileData{}

	query := "SELECT * FROM files WHERE user_id=$1"
	rows, err := f.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		fd := &model.FileData{}
		if err := rows.Scan(&fd.Id, &fd.UserId, &fd.FileName, &fd.Size, &fd.UploadDate); err != nil {
			return nil, err
		}

		files = append(files, fd)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return files, nil
}

func (f *FilesPostgres) DeleteFile(fileId, userId uint64) error {
	query := "DELETE FROM files WHERE user_id=$1 AND file_id=$2"
	_, err := f.db.Exec(query, userId, fileId)

	return err
}
