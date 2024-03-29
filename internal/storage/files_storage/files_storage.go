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

func (f *FilesMinio) UploadFile(ctx context.Context, fd *model.FileData, file io.Reader) error {
	objectName := fmt.Sprintf("%d/%s", fd.UserId, fd.FileName)

	_, err := f.client.minioClient.PutObject(ctx, f.client.bucketName, objectName, file, fd.Size, minio.PutObjectOptions{})
	return err
}

func (f *FilesMinio) GetFile(ctx context.Context, userId uint64, fd *model.FileData) ([]byte, error) {
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
