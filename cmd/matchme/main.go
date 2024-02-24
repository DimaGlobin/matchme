package main

import (
	"os"

	"golang.org/x/exp/slog"

	"github.com/DimaGlobin/matchme/internal/config"
)

/*
TODO: init config
TODO: init logger
*/

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// fmt.Println(os.Getwd())
	cfg := config.MustLoad("config/server_config.yaml")
	log := setupLogger(cfg.Env)
	log.Info("starting application", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log

}
