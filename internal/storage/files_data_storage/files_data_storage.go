package files_data_storage

import (
	"github.com/DimaGlobin/matchme/internal/model"
	"gorm.io/gorm"
)

type FilesPostgres struct {
	db *gorm.DB
}

func NewFilesPostgres(db *gorm.DB) *FilesPostgres {
	return &FilesPostgres{db: db}
}

func (f *FilesPostgres) AddFile(data *model.FileData) (uint64, error) {
	if err := f.db.Create(data).Error; err != nil {
		return 0, err
	}
	return data.Id, nil
}

func (f *FilesPostgres) GetFilesCount(userId uint64) (int64, error) {
	var count int64
	if err := f.db.Model(&model.FileData{}).Where("user_id = ?", userId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (f *FilesPostgres) GetFileById(fileId, userId uint64) (*model.FileData, error) {
	fd := &model.FileData{}
	if err := f.db.Where("file_id = ? AND user_id = ?", fileId, userId).First(fd).Error; err != nil {
		return nil, err
	}
	return fd, nil
}

func (f *FilesPostgres) GetFileByName(userId uint64, filename string) (*model.FileData, error) {
	fd := &model.FileData{}
	if err := f.db.Where("file_name = ? AND user_id = ?", filename, userId).First(fd).Error; err != nil {
		return nil, err
	}
	return fd, nil
}

func (f *FilesPostgres) GetAllFiles(userId uint64) ([]*model.FileData, error) {
	var files []*model.FileData
	if err := f.db.Where("user_id = ?", userId).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func (f *FilesPostgres) DeleteFile(fileId, userId uint64) error {
	if err := f.db.Where("file_id = ? AND user_id = ?", fileId, userId).Delete(&model.FileData{}).Error; err != nil {
		return err
	}
	return nil
}
