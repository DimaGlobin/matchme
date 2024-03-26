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
	objectName := fmt.Sprintf("%d/%s", fd.Id, fd.FileName)

	_, err := f.client.minioClient.PutObject(ctx, f.client.bucketName, objectName, file, fd.Size, minio.PutObjectOptions{})
	return err
}
