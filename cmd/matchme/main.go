package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/exp/slog"

	"github.com/DimaGlobin/matchme/internal/config"
	logger "github.com/DimaGlobin/matchme/internal/lib/logger/common"
	"github.com/DimaGlobin/matchme/internal/lib/logger/sl"
	"github.com/DimaGlobin/matchme/internal/server"
	storage "github.com/DimaGlobin/matchme/internal/storage/users_storage"
	"github.com/go-chi/chi/v5"
)

/*
TODO: init config
TODO: init logger
*/

func main() {
	// fmt.Println(os.Getwd())
	cfg := config.MustLoad("./config/server_config.yaml")
	log := logger.NewCommonLogger(cfg.Env)

	log.Info(
		"starting url-shortener",
		slog.String("env", cfg.Env),
		slog.String("version", "123"),
	)
	log.Debug("debug messages are enabled")

	_, err := storage.NewConnection(&cfg.UsersDBConfig)
	if err != nil {
		log.Error("Cannot connect to users db")
		return
	}

	log.Info("starting server", slog.String("address", cfg.Address))

	srv := new(server.Server)
	router := chi.NewRouter()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err = srv.Run(cfg, router); err != nil {
			log.Error("Cannot run server")
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
