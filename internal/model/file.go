package model

import "time"

type FileData struct {
	Id         uint64    `db:"file_id"`
	UserId     uint64    `db:"user_id"`
	FileName   string    `db:"file_name"`
	Size       int64     `db:"size"`
	UploadDate time.Time `db:"upload_date"`
}

type File struct {
	Name   string
	Buffer []byte
}
