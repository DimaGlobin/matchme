package files_handlers

import (
	"net/http"

	"github.com/DimaGlobin/matchme/internal/service"
	"golang.org/x/exp/slog"
)

type GetFileHandler struct {
	logger  *slog.Logger
	service *service.Service
}

func NewGetFileHander(log *slog.Logger, srv *service.Service) *GetFileHandler {
	return &GetFileHandler{
		logger:  log,
		service: srv,
	}
}

func (g *GetFileHandler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
