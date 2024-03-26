package model

type FileData struct {
	Id       int    `db:"file_id"`
	UserId   int    `db:"user_id"`
	FileName string `db:"file_name"`
	Size     int64  `db:"size"`
}
