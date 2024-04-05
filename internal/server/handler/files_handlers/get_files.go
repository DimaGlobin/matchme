package files_handlers

import (
	"mime/multipart"
	"net/http"

	"github.com/DimaGlobin/matchme/internal/lib/api"
	"github.com/DimaGlobin/matchme/internal/lib/logger/sl"
	"github.com/DimaGlobin/matchme/internal/middleware/auth"
	"github.com/DimaGlobin/matchme/internal/service"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
)

type GetFilesHandler struct {
	logger  *slog.Logger
	service *service.Service
}

func NewGetFilesHandler(log *slog.Logger, srv *service.Service) *GetFilesHandler {
	return &GetFilesHandler{
		logger:  log,
		service: srv,
	}
}

func (g *GetFilesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		log := g.logger.With(
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		userId := r.Context().Value(auth.UserCtx).(uint64)

		writer := multipart.NewWriter(w)
		w.Header().Set("Content-Type", writer.FormDataContentType())

		files, err := g.service.GetAllFiles(r.Context(), userId)
		if err != nil {
			msg := "Unable to get files"
			log.Error(msg, sl.Err(err))
			api.Respond(w, r, http.StatusInternalServerError, msg)

			return
		}

		for _, file := range files {
			part, err := writer.CreateFormFile("files", file.Name)
			if err != nil {
				msg := "Unable to create form file"
				log.Error(msg, sl.Err(err))
				api.Respond(w, r, http.StatusInternalServerError, msg)

				return
			}

			_, err = part.Write(file.Buffer)
			if err != nil {
				msg := "Unable to write files"
				log.Error(msg, sl.Err(err))
				api.Respond(w, r, http.StatusInternalServerError, msg)

				return
			}
		}

		msg := "Files was successfully sent"
		log.Info(msg)
		api.Respond(w, r, http.StatusOK, "")
	}
