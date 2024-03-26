package files_handlers

import (
	"net/http"

	"github.com/DimaGlobin/matchme/internal/service"
	"golang.org/x/exp/slog"
)

type DeleteFileHandler struct {
	logger  *slog.Logger
	service *service.Service
}

func NewDeleteFileHandler(log *slog.Logger, srv *service.Service) *DeleteFileHandler {
	return &DeleteFileHandler{
		logger:  log,
		service: srv,
	}
}

func (d *DeleteFileHandler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
