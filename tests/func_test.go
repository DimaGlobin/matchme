package tests

import (
	"fmt"
	"net/http/httptest"
	"os/exec"

	stdlog "log"
	"testing"

	"github.com/DimaGlobin/matchme/internal/config"
	logger "github.com/DimaGlobin/matchme/internal/lib/logger/common"
	"github.com/DimaGlobin/matchme/internal/server"
	"github.com/DimaGlobin/matchme/internal/server/router"
	"github.com/DimaGlobin/matchme/internal/service"
	"github.com/DimaGlobin/matchme/internal/storage"
	"github.com/DimaGlobin/matchme/internal/storage/cache_storage"
	"github.com/DimaGlobin/matchme/internal/storage/files_storage"
	"github.com/lamoda/gonkey/runner"
	"golang.org/x/exp/slog"
)

func Setup() *httptest.Server {
	cfg := &config.Config{}
	cfg = config.MustLoad("../config/server_config.yaml")

	log := logger.NewCommonLogger(cfg.Env)

	db, err := storage.NewPostgresDB(cfg)
	if err != nil {
		fmt.Println(cfg, err)
		stdlog.Fatal("Cannot connect to db")
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", cfg.UsersDBConfig.User, cfg.UsersDBConfig.Password, cfg.UsersDBConfig.Host, cfg.UsersDBConfig.Port, cfg.UsersDBConfig.DBName, cfg.UsersDBConfig.SSLMode)
	fmt.Println(dbURL)
	cmd := exec.Command("migrate", "-path", "../internal/schema/migrations", "-database", dbURL, "up")
	err = cmd.Run()
	if err != nil {
		fmt.Println(cfg, err.Error())
		stdlog.Fatal("Cannot migrate DB")
	}

	minioClient, err := files_storage.NewMinioClient(cfg)
	if err != nil {
		fmt.Println(cfg, err)
		stdlog.Fatal("Cannot connect to files storage")
	}

	rdb, err := cache_storage.NewRedisDB(cfg)
	if err != nil {
		fmt.Println(cfg, err)
		stdlog.Fatal("Cannot connect to cache storage")
	}

	storage := storage.NewStorage(db, minioClient, rdb)
	service := service.NewService(*storage)

	log.Info(
		"starting url-shortener",
		slog.String("env", cfg.Env),
		slog.String("version", "123"),
	)
	log.Debug("debug messages are enabled")

	log.Info("starting server", slog.String("address", cfg.Address))

	router := router.NewRouter(log, service)
	server := server.NewHTTPServer(cfg, router)

	srv := httptest.NewServer(server.HttpServer.Handler)

	return srv
}

func RunTest(t *testing.T, testPath string) {
	srv := Setup()

	runner.RunWithTesting(t, &runner.RunWithTestingParams{
		Server:   srv,
		TestsDir: testPath,
	})
}

func TestApi(t *testing.T) {
	RunTest(t, "./cases/users_cases/sign_up_cases.yaml")
}
