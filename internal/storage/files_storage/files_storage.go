package files_storage

import (
	"context"
	"fmt"
	"io"

	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/minio/minio-go/v7"
)

type FilesMinio struct {
	client *MinioClient
}

func NewFilesMinio(client *MinioClient) *FilesMinio {
	return &FilesMinio{
		client: client,
	}
}

func (f *FilesMinio) UploadFile(fd *model.FileData, file io.Reader) error {
	ctx := context.TODO()
	objectName := fmt.Sprintf("%d/%s", fd.UserId, fd.FileName)

	_, err := f.client.minioClient.PutObject(ctx, f.client.bucketName, objectName, file, fd.Size, minio.PutObjectOptions{})
	return err
}

func (f *FilesMinio) GetFile(userId uint64, fd *model.FileData) ([]byte, error) {
	ctx := context.TODO()
	url := fmt.Sprintf("%d/%s", userId, fd.FileName)
	obj, err := f.client.minioClient.GetObject(ctx, f.client.bucketName, url, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer obj.Close()

	objInfo, err := obj.Stat()
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, objInfo.Size)

	_, err = obj.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return buffer, nil
}

func (f *FilesMinio) DeleteFile(userId uint64, fd *model.FileData) error {
	ctx := context.TODO()
	url := fmt.Sprintf("%d/%s", userId, fd.FileName)

	return f.client.minioClient.RemoveObject(ctx, f.client.bucketName, url, minio.RemoveObjectOptions{})
}

// func (f *FilesMinio) GetAllFiles(ctx context.Context, userId uint64) ([][]byte, error) {

// }
