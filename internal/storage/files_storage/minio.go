package files_storage

import (
	"context"

	"github.com/DimaGlobin/matchme/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	minioClient *minio.Client
	bucketName  string
}

func NewMinioClient(cfg *config.Config) (*MinioClient, error) {
	minioClient, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.ScretKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	exists, err := minioClient.BucketExists(context.Background(), cfg.Bucketname)
	if err != nil {
		return nil, err
	}

	if !exists {
		err = minioClient.MakeBucket(context.Background(), cfg.Bucketname, minio.MakeBucketOptions{})
		if err != nil {
			return nil, err
		}
	}

	return &MinioClient{minioClient: minioClient, bucketName: cfg.Bucketname}, nil
}
