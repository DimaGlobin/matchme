package users_handler

import (
	"net/http"

	"github.com/DimaGlobin/matchme/internal/lib/api"
	"github.com/DimaGlobin/matchme/internal/lib/logger/sl"
	"github.com/DimaGlobin/matchme/internal/service"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
)

type SigninHandler struct {
	logger  *slog.Logger
	service *service.Service
}

func NewSignInHandler(log *slog.Logger, srv *service.Service) *SigninHandler {
	return &SigninHandler{
		logger:  log,
		service: srv,
	}
}

type SignInBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *SigninHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body := &SignInBody{}
	log := s.logger.With(
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	err := render.DecodeJSON(r.Body, body)
	if err != nil {
		msg := "Unable to decode request body"
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusBadRequest, msg)

		return
	}

	token, err := s.service.UsersService.GenerateToken(body.Email, body.Password)
	if err != nil {
		msg := "Unable to create jwt token"
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusInternalServerError, msg)

		return
	}

	api.Respond(w, r, http.StatusOK, map[string]interface{}{
		"token": token,
	})

}
