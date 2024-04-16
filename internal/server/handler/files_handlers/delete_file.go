package files_handlers

import (
	"net/http"
	"strconv"

	"github.com/DimaGlobin/matchme/internal/lib/api"
	"github.com/DimaGlobin/matchme/internal/lib/logger/sl"
	"github.com/DimaGlobin/matchme/internal/middleware/auth"
	"github.com/DimaGlobin/matchme/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

func (d *DeleteFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := d.logger.With(
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userId := r.Context().Value(auth.UserCtx).(uint64)

	fileIdStr := chi.URLParam(r, "id")
	fileId, err := strconv.ParseUint(fileIdStr, 10, 64)
	if err != nil {
		msg := "Cannot parse url query"
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusBadRequest, "")

		return
	}

	if err := d.service.FilesServiceInt.DeleteFile(fileId, userId); err != nil {
		msg := "Cannot delete file"
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusInternalServerError, msg)

		return
	}

	msg := "File was deleted successfully"
	log.Info(msg)
	api.Respond(w, r, http.StatusOK, msg)
}
