package main

import (
	"context"
	"fmt"
	stdlog "log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/exp/slog"

	"github.com/DimaGlobin/matchme/internal/config"
	logger "github.com/DimaGlobin/matchme/internal/lib/logger/common"
	"github.com/DimaGlobin/matchme/internal/lib/logger/sl"
	"github.com/DimaGlobin/matchme/internal/server"
	"github.com/DimaGlobin/matchme/internal/server/router"
	"github.com/DimaGlobin/matchme/internal/service"
	"github.com/DimaGlobin/matchme/internal/storage"
	"github.com/DimaGlobin/matchme/internal/storage/cache_storage"
	"github.com/DimaGlobin/matchme/internal/storage/files_storage"
	"github.com/joho/godotenv"
)

/*
TODO: init config
TODO: init logger
*/

// @title MatchMe API
// @version 1.0
// @description API Server for MatchMe application

// @host localhost:8084
// @basePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// fmt.Println(os.Getwd())
	if err := godotenv.Load("./config/.env"); err != nil {
		stdlog.Fatal("Cannot load .env file")
	}

	cfg := &config.Config{}

	env := os.Getenv("ENV")
	if env == "test" {
		cfg = config.MustLoad("./tests/test_env/test_config.yaml")
	} else {
		cfg = config.MustLoad("./config/server_config.yaml")
	}

	log := logger.NewCommonLogger(cfg.Env)

	db, err := storage.NewPostgresDB(cfg)
	if err != nil {
		fmt.Println(cfg, err)
		stdlog.Fatal("Cannot connect to db")
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", cfg.UsersDBConfig.User, cfg.UsersDBConfig.Password, cfg.UsersDBConfig.Host, cfg.UsersDBConfig.Port, cfg.UsersDBConfig.DBName, cfg.UsersDBConfig.SSLMode)

	cmd := exec.Command("migrate", "-path", "./internal/schema/migrations", "-database", dbURL, "up")
	err = cmd.Run()
	if err != nil {
		fmt.Println(cfg, err)
		stdlog.Fatal("Cannot migrate DB")
		return
	}

	minioClient, err := files_storage.NewMinioClient(cfg)
	if err != nil {
		fmt.Println(cfg, err)
		stdlog.Fatal("Cannot connect to files storage")
		return
	}

	rdb, err := cache_storage.NewRedisDB(cfg)
	if err != nil {
		fmt.Println(cfg, err)
		stdlog.Fatal("Cannot connect to cache storage")
		return
	}

	fmt.Printf("rdb: %p\n", rdb)

	storage := storage.NewStorage(db, minioClient, rdb)
	fmt.Println("storage: ", *storage)
	service := service.NewService(*storage)
	fmt.Println("service: ", service)

	log.Info(
		"starting matchme",
		slog.String("env", cfg.Env),
		slog.String("version", "123"),
	)
	log.Debug("debug messages are enabled")

	// _, err := storage.NewConnection(&cfg.UsersDBConfig)
	// if err != nil {
	// 	log.Error("Cannot connect to users db")
	// 	return
	// }

	log.Info("starting server", slog.String("address", cfg.Address))

	router := router.NewRouter(log, service)
	srv := server.NewHTTPServer(cfg, router)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.Run(); err != nil {
			log.Error("Cannot run server", sl.Err(err))
			return
		}
	}()

	log.Info("server started")

	<-done
	log.Info("stopping server")

	// TODO: move timeout to config
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", sl.Err(err))

		return
	}

	// TODO: close storage

	log.Info("server stopped")
}
