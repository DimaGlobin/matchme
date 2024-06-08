package files_handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/DimaGlobin/matchme/internal/lib/api"
	"github.com/DimaGlobin/matchme/internal/lib/logger/sl"
	"github.com/DimaGlobin/matchme/internal/middleware/auth"
	"github.com/DimaGlobin/matchme/internal/service"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
)

type GetFileByNameHandler struct {
	logger  *slog.Logger
	service *service.Service
}

func NewGetFileByNameHandler(log *slog.Logger, srv *service.Service) *GetFileByNameHandler {
	return &GetFileByNameHandler{
		logger:  log,
		service: srv,
	}
}

// @Summary GetPhotoByName
// @Security BearerAuth
// @Tags api
// @Description Get user's photo and send like multipart/form-data
// @ID get-photo-by-name
// @Produce multipart/form-data
// @Success 200 {object} string "File was successfully sent"
// @Failure 400 {string} string "Empty file name"
// @Failure 500
// @Router /api/photos/filename/{filename} [get]
func (g *GetFileByNameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := g.logger.With(
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	parts := strings.Split(r.URL.Path, "/") // I understand that it's not the best solution but
	filename := parts[len(parts)-1]         // I'm tired of searching how to do it using built in
	// chi tools :( yours CEO
	if filename == "" {
		msg := "Empty file name"
		api.Respond(w, r, http.StatusBadRequest, msg)
		return
	}
	// filename := chi.URLParam(r, "filename")
	fmt.Println(filename)

	userId := r.Context().Value(auth.UserIdKey).(uint64)

	file, err := g.service.FilesServiceInt.GetFileByName(userId, filename)
	if err != nil {
		msg := "Unable to get file"
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusInternalServerError, msg)

		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Name))

	_, err = w.Write(file.Buffer)
	if err != nil {
		msg := "Unable to send file in response"
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusInternalServerError, msg)

		return
	}

	msg := "File was successfully sent"
	log.Info(msg)
	api.Respond(w, r, http.StatusOK, msg)
}
