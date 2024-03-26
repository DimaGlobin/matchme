package files_handlers

import (
	"net/http"

	"github.com/DimaGlobin/matchme/internal/lib/api"
	"github.com/DimaGlobin/matchme/internal/lib/logger/sl"
	"github.com/DimaGlobin/matchme/internal/middleware/auth"
	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/service"
	"golang.org/x/exp/slog"
)

type UploadFileHandler struct {
	logger  *slog.Logger
	service *service.Service
}

func NewUploadFileHandler(log *slog.Logger, srv *service.Service) *UploadFileHandler {
	return &UploadFileHandler{
		logger:  log,
		service: srv,
	}
}

func (u *UploadFileHandler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// for name, values := range r.Header {
		// 	for _, value := range values {
		// 		fmt.Println(name, value)
		// 	}
		// }

		reqFile, header, err := r.FormFile("file")
		if err != nil {
			msg := "Cannot get file from request"
			u.logger.Error(msg, sl.Err(err))
			api.Respond(w, r, http.StatusBadRequest, "")

			return
		}
		defer reqFile.Close()

		fileData := new(model.FileData)
		fileData.FileName = header.Filename
		fileData.Size = header.Size
		fileData.UserId = r.Context().Value(auth.UserCtx).(int)

		err = u.service.FilesService.UploadFile(r.Context(), fileData, reqFile)
		if err != nil {
			msg := "Cannot upload file"
			u.logger.Error(msg, sl.Err(err))
			api.Respond(w, r, http.StatusInternalServerError, "")

			return
		}

		msg := "File was successfully uploaded"
		u.logger.Info(msg)
		api.Respond(w, r, http.StatusOK, msg)
	}
}
