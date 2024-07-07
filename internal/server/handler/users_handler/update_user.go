package users_handler

import (
	"net/http"

	"github.com/DimaGlobin/matchme/internal/lib/api"
	"github.com/DimaGlobin/matchme/internal/lib/logger/sl"
	"github.com/DimaGlobin/matchme/internal/middleware/auth"
	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/service"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
)

type UpdateUserHandler struct {
	logger  *slog.Logger
	service *service.Service
}

func NewUpdateUserHandler(log *slog.Logger, srv *service.Service) *UpdateUserHandler {
	return &UpdateUserHandler{
		logger:  log,
		service: srv,
	}
}

// @Summary UpdateUser
// @Security BearerAuth
// @Tags api
// @Description update user info
// @ID update-user
// @Accept  json
// @Produce  json
// @Param input body model.Updates true "User updates"
// @Success 200
// @Failure 400,401
// @Failure 500
// @Router /api/users/ [put]
func (s *UpdateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := s.logger.With(
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var updates model.Updates

	err := render.DecodeJSON(r.Body, &updates)
	if err != nil {
		msg := "Unable to decode request body"
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusBadRequest, msg)

		return
	}

	// fmt.Println("Updates: ", updates)

	if err := updates.Valid(); err != nil {
		msg := "Invalid request body"
		log.Error(msg)
		api.Respond(w, r, http.StatusBadRequest, msg)

		return
	}

	// fmt.Println(updates)

	user_id := r.Context().Value(auth.UserIdKey).(uint64)
	err = s.service.UsersServiceInt.UpdateUser(user_id, updates)
	if err != nil {
		msg := "Cannot update user"
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusInternalServerError, msg)

		return
	}

	msg := "User was successfully updated"
	log.Info(msg)
	api.Respond(w, r, http.StatusOK, msg)
}
