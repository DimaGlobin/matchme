package files_handlers

import (
	"net/http"
	"strings"

	"github.com/DimaGlobin/matchme/internal/lib/api"
	"github.com/DimaGlobin/matchme/internal/lib/logger/sl"
	"github.com/DimaGlobin/matchme/internal/middleware/auth"
	"github.com/DimaGlobin/matchme/internal/service"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
)

//  а обязательно ли везде передавать логгер, по-моему у него есть свойство быть глобальным!!!
// если мы проинициализируем его всего один раз, могу ошибаться, поправь, если это не так

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

// @Summary DeleteFileById
// @Security BearerAuth
// @Tags api
// @Description Delete user's photo
// @ID delete-photo-by-id
// @Success 200 {object} string "File was deleted successfully"
// @Failure 400 {string} string "Empty file name"
// @Failure 500 {string} string "Cannot parse url query"
// @Router /api/photos/{filename} [delete]
func (d *DeleteFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := d.logger.With(
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userId := r.Context().Value(auth.UserIdKey).(uint64)

	parts := strings.Split(r.URL.Path, "/") // I understand that it's not the best solution but
	filename := parts[len(parts)-1]         // I'm tired of searching how to do it using built in
	// chi tools :( yours CEO
	if filename == "" {
		msg := "Empty file name"
		api.Respond(w, r, http.StatusBadRequest, msg)
		return
	}
	// filename := chi.URLParam(r, "filename")
	// fmt.Println(filename)

	if err := d.service.FilesServiceInt.DeleteFile(userId, filename); err != nil {
		msg := "Cannot delete file"
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusInternalServerError, msg)

		return
	}

	msg := "File was deleted successfully"
	log.Info(msg)
	api.Respond(w, r, http.StatusOK, msg)
}
