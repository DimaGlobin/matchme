package main

import (
	"context"
	stdlog "log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/exp/slog"

	"github.com/DimaGlobin/matchme/internal/config"
	logger "github.com/DimaGlobin/matchme/internal/lib/logger/common"
	"github.com/DimaGlobin/matchme/internal/lib/logger/sl"
	"github.com/DimaGlobin/matchme/internal/server"
	"github.com/go-chi/chi/v5"
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

	router := chi.NewRouter()
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
