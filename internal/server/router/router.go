package router

import (
	"net/http"

	"github.com/DimaGlobin/matchme/internal/middleware/auth"
	"github.com/DimaGlobin/matchme/internal/middleware/logger"
	"github.com/DimaGlobin/matchme/internal/server/handler/users_handler"
	"github.com/DimaGlobin/matchme/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
)

func NewRouter(log *slog.Logger, srv *service.Service) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(logger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	router.Post("/sign_up", users_handler.NewSignUpHandler(log, srv))
	router.Post("/sign_in", users_handler.NewSignInHandler(log, srv))

	router.Mount("/api", NewApiRouter(log, srv))

	return router
}

func NewApiRouter(log *slog.Logger, srv *service.Service) chi.Router {
	router := chi.NewRouter()

	router.Use(auth.New(log, srv))

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	return router
}
