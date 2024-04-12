package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env                string `yaml:"env" env-default:"local"`
	HTTPServer         `yaml:"http_server"`
	UsersDBConfig      `yaml:"users_db"`
	FilesStorageConfig `yaml:"files_storage"`
	CacheStorage       `yaml:"cache_storage"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type UsersDBConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Port     int    `yaml:"port"`
	SSLMode  string `yaml:"sslmode"`
}

type FilesStorageConfig struct {
	Endpoint   string `yaml:"endpoint"`
	AccessKey  string `yaml:"access_key"`
	ScretKey   string `yaml:"secret_key"`
	Bucketname string `yaml:"bucket_name"`
}

type CacheStorage struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func MustLoad(configPath string) *Config {

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("config path doesn't exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Cannot read config: %s", configPath)
	}

	// dir, _ := os.Getwd()
	// fmt.Println("Config: ", cfg, "\nConfigPath: ", configPath, "\nPWD: ", dir)

	return &cfg
}
