package model

import "time"

type FileData struct {
	Id         uint64    `gorm:"primaryKey;column:file_id"`
	UserId     uint64    `gorm:"column:user_id"`
	FileName   string    `gorm:"column:file_name"`
	Size       int64     `gorm:"column:size"`
	UploadDate time.Time `gorm:"column:upload_date"`
}

func (FileData) TableName() string {
	return "files"
}

type File struct {
	Name   string
	Buffer []byte // публичное поле?
}
