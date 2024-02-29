package storage

import (
	"fmt"

	"github.com/DimaGlobin/matchme/internal/config"
	"github.com/DimaGlobin/matchme/internal/storage/users_storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg *config.UsersDBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&users_storage.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
