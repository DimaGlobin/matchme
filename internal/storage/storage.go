package storage

import (
	"github.com/DimaGlobin/matchme/internal/storage/cache_storage"
	"github.com/DimaGlobin/matchme/internal/storage/files_data_storage"
	"github.com/DimaGlobin/matchme/internal/storage/files_storage"
	"github.com/DimaGlobin/matchme/internal/storage/ratings_storage"
	"github.com/DimaGlobin/matchme/internal/storage/users_storage"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Storage struct {
	users_storage.UsersStorage
	files_storage.FilesStorage
	files_data_storage.FilesDataStorage
	ratings_storage.RatingsStorage
	cache_storage.CacheStorage
}

func NewStorage(db *sqlx.DB, minioClient *files_storage.MinioClient, rdb *redis.Client) *Storage {
	return &Storage{
		UsersStorage:     users_storage.NewUsersPostgres(db),
		FilesStorage:     files_storage.NewFilesMinio(minioClient),
		FilesDataStorage: files_data_storage.NewFilesPostgres(db),
		RatingsStorage:   ratings_storage.NewRatingsPostgres(db),
		CacheStorage:     cache_storage.NewCacheRedis(rdb),
	}
}
