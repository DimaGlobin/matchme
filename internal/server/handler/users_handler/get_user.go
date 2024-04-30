package users_handler

import (
	"net/http"

	"github.com/DimaGlobin/matchme/internal/lib/api"
	"github.com/DimaGlobin/matchme/internal/lib/logger/sl"
	"github.com/DimaGlobin/matchme/internal/middleware/auth"
	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/service"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
)

type GetUserHandler struct {
	logger  *slog.Logger
	service *service.Service
}

func NewGetUserHandler(log *slog.Logger, srv *service.Service) *GetUserHandler {
	return &GetUserHandler{
		logger:  log,
		service: srv,
	}
}

// @Summary GetUser
// @Security BearerAuth
// @Tags api
// @Description get user
// @ID get-user
// @Accept  json
// @Produce  json
// @Success 200 {object} model.UserInfo
// @Failure 400,401
// @Failure 500
// @Router /api/users/ [get]
func (s *GetUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := s.logger.With(
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	id := r.Context().Value(auth.UserCtx).(uint64)

	user, err := s.service.UsersServiceInt.GetuserById(id)
	if err != nil {
		msg := "Unable to get user"
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusInternalServerError, msg)

		return
	}

	api.Respond(w, r, http.StatusOK, map[string]interface{}{
		"user": model.UserInfo{
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Name:        user.Name,
			Sex:         user.Sex,
			Age:         user.Age,
			City:        user.City,
			Description: user.Description,
			MaxAge:      user.MaxAge,
		},
	})
}
