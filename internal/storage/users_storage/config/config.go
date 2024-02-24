package storage

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type UsersDBConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Port     int    `yaml:"port"`
	SSLMode  string `yaml:"sslmode"`
}

func MustLoad() *UsersDBConfig {
	configPath := "config.yaml"

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("UsesrDB config path doesn't exist: %s", configPath)
	}

	var cfg UsersDBConfig

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Cannot read UsersDB config: %s", configPath)
	}

	return &cfg
}
