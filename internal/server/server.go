package server

import (
	"context"
	"net/http"

	"github.com/DimaGlobin/matchme/internal/config"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	httpServer *http.Server
}

func NewHTTPServer(cfg *config.Config, router chi.Router) *Server {
	var s Server

	s.httpServer = &http.Server{
		Addr:           cfg.Address,
		Handler:        router,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    cfg.Timeout,
		WriteTimeout:   cfg.Timeout,
		IdleTimeout:    cfg.IdleTimeout,
	}

	return &s
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
