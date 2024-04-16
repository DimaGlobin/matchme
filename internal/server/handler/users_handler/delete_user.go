package users_handler

import (
	"net/http"

	"github.com/DimaGlobin/matchme/internal/lib/api"
	"github.com/DimaGlobin/matchme/internal/lib/logger/sl"
	"github.com/DimaGlobin/matchme/internal/middleware/auth"
	"github.com/DimaGlobin/matchme/internal/service"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
)

type DeleteUserHandler struct {
	logger  *slog.Logger
	service *service.Service
}

func NewDeleteUserHandler(log *slog.Logger, srv *service.Service) *DeleteUserHandler {
	return &DeleteUserHandler{
		logger:  log,
		service: srv,
	}
}

func (s *DeleteUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := s.logger.With(
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	user_id := r.Context().Value(auth.UserCtx).(uint64)
	err := s.service.UsersServiceInt.DeleteUser(user_id)
	if err != nil {
		msg := "Cannot delete user"
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusInternalServerError, msg)

		return
	}

	msg := "User was successfully deleted"
	log.Info(msg)
	api.Respond(w, r, http.StatusOK, msg)
}
