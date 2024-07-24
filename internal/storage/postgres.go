package storage

import (
	"fmt"

	"github.com/DimaGlobin/matchme/internal/config"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// cfg -> MustLoad, ; potsgresql -> MustLoad;

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.UsersDBConfig.Host, cfg.UsersDBConfig.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// может сразу упадём, если всё плохо?

	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
