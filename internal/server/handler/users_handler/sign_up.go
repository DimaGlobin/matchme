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

type idResponse struct {
	Id uint64 `json:"id"`
}

func NewSignUpHandler(log *slog.Logger, srv *service.Service) *SignUpHandler {
	return &SignUpHandler{
		logger:  log,
		service: srv,
	}
}

// @Summary SignUp
// @Tags auth
// @Description signup
// @ID sign-up
// @Accept  json
// @Produce  json
// @Param input body model.User true "User information"
// @Success 200 {object} idResponse "idResponse"
// @Failure 400
// @Failure 500
// @Router /auth/sign_up [post]
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

	id, err := s.service.UsersServiceInt.CreateUser(user)
	if err != nil {
		msg := err.Error()
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusInternalServerError, err.Error())

		return
	}

	log.Info(fmt.Sprintf("User was successfully created, id: %d", id))
	api.Respond(w, r, http.StatusOK, idResponse{
		Id: id,
	})

}
