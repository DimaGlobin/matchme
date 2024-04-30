package files_handlers

import (
	"fmt"
	"net/http"

	"github.com/DimaGlobin/matchme/internal/lib/api"
	"github.com/DimaGlobin/matchme/internal/lib/logger/sl"
	"github.com/DimaGlobin/matchme/internal/middleware/auth"
	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/service"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
)

type PhotoIdResponse struct {
	Id uint64 `json:"id"`
}

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

// @Summary UploadPhoto
// @Security BearerAuth
// @Tags api
// @Description upload photo
// @ID upload-photo
// @Accept  json
// @Produce  json
// @Param file formData file true "Photo to upload"
// @Success 200 {object} PhotoIdResponse
// @Failure 500
// @Router /api/photos/ [post]
func (u *UploadFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// for name, values := range r.Header {
	// 	for _, value := range values {
	// 		fmt.Println(name, value)
	// 	}
	// }

	log := u.logger.With(
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

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
	fileData.UserId = r.Context().Value(auth.UserCtx).(uint64)

	if fileData.Size > 3*1024*1024 {
		msg := "File size shiuld be less than 3 MB"
		api.Respond(w, r, http.StatusBadRequest, msg)

		return
	}

	id, err := u.service.FilesServiceInt.UploadFile(fileData, reqFile)
	if err != nil {
		msg := "Cannot upload file"
		u.logger.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusInternalServerError, "")

		return
	}

	log.Info(fmt.Sprintf("User was successfully created, id: %d", id))
	api.Respond(w, r, http.StatusOK, PhotoIdResponse{
		Id: id,
	})
}
