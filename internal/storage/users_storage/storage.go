package storage

import (
	"fmt"

	"github.com/DimaGlobin/matchme/internal/config"
	models "github.com/DimaGlobin/matchme/internal/storage/users_storage/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(cfg *config.UsersDBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewUsersDB(cfg *config.UsersDBConfig) (*gorm.DB, error) {
	DB, err := NewConnection(cfg)
	if err != nil {
		return nil, err
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	return DB, nil
}
