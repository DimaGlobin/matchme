package users_handler

import (
	"fmt"
	"net/http"

	"github.com/DimaGlobin/matchme/internal/lib/api"
	"github.com/DimaGlobin/matchme/internal/lib/logger/sl"
	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/service"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
)

type SignUpHandler struct {
	logger  *slog.Logger
	service *service.Service
}

func NewSignUpHandler(log *slog.Logger, srv *service.Service) *SignUpHandler {
	return &SignUpHandler{
		logger:  log,
		service: srv,
	}
}

func (s *SignUpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}

	log := s.logger.With(
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	err := render.DecodeJSON(r.Body, user)
	if err != nil {
		msg := "Unable to decode request body"
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusBadRequest, msg)

		return
	}

	id, err := s.service.UsersService.CreateUser(user)
	if err != nil {
		msg := "Unable to create user"
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusInternalServerError, "")

		return
	}

	log.Info(fmt.Sprintf("User was successfully created, id: %d", id))
	api.Respond(w, r, http.StatusOK, map[string]interface{}{
		"id": id,
	})

}
