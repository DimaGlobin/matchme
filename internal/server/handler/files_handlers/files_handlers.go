package files_handlers

import (
	"net/http"

	"github.com/DimaGlobin/matchme/internal/service"
	"golang.org/x/exp/slog"
)

func NewUploadFileHandler(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func NewDeleteFileHandler(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func NewGetFileHandler(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
