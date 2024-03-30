package files_handlers

import (
	"fmt"
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

type GetFileByIdHandler struct {
	logger  *slog.Logger
	service *service.Service
}

func NewGetFileByIdHander(log *slog.Logger, srv *service.Service) *GetFileByIdHandler {
	return &GetFileByIdHandler{
		logger:  log,
		service: srv,
	}
}

func (g *GetFileByIdHandler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := g.logger.With(
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		fileIdStr := chi.URLParam(r, "id")
		fmt.Println("idStr: ", fileIdStr)
		fileId, err := strconv.ParseUint(fileIdStr, 10, 64)
		if err != nil {
			msg := "Unable to get id from url query"
			log.Error(msg, sl.Err(err))
			api.Respond(w, r, http.StatusBadRequest, msg)

			return
		}

		userId := r.Context().Value(auth.UserCtx).(uint64)

		file, err := g.service.GetFileById(r.Context(), fileId, userId)
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
}
