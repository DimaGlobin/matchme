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
	"github.com/joho/godotenv"
)

/*
TODO: init config
TODO: init logger
*/

func main() {
	// fmt.Println(os.Getwd())
	if err := godotenv.Load("./config/.env"); err != nil {
		stdlog.Fatal("Cannot load .env file")
	}

	cfg := config.MustLoad("./config/server_config.yaml")
	log := logger.NewCommonLogger(cfg.Env)

	db, err := storage.NewPostgresDB(cfg)
	if err != nil {
		fmt.Println(cfg, err)
		stdlog.Fatal("Cannot connect to db")
	}

	cmd := exec.Command("migrate", "-path", "./internal/schema/migrations", "-database", "postgres://postgres:qwerty@db:5432/postgres?sslmode=disable", "up")
	err = cmd.Run()
	if err != nil {
		stdlog.Fatal("Cannot migrate DB")
		return
	}

	storage := storage.NewStorage(db)
	service := service.NewService(*storage)

	log.Info(
		"starting url-shortener",
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
