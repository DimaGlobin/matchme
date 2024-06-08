package users_handler

import (
	"net/http"

	"github.com/DimaGlobin/matchme/internal/lib/api"
	"github.com/DimaGlobin/matchme/internal/lib/logger/sl"
	"github.com/DimaGlobin/matchme/internal/mm_errors"
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

// @Summary SignUp
// @Tags auth
// @Description signup
// @ID sign-up
// @Accept  json
// @Produce  json
// @Param input body model.User true "User information"
// @Success 200 {object} api.IdResponse "IdResponse"
// @Failure 400 {object} api.ErrResponse "ErrResponse"
// @Failure 500 {object} api.ErrResponse "ErrResponse"
// @Router /auth/sign_up [post]
func (s *SignUpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}

	log := s.logger.With(
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	err := render.DecodeJSON(r.Body, user)
	if err != nil {
		log.Error(mm_errors.DecodeError.Error(), sl.Err(err))
		api.Respond(w, r, http.StatusBadRequest, api.ErrResponse{
			Err: mm_errors.DecodeError.Error(),
		})

		return
	}

	if err = user.Valid(); err != nil {
		log.Error(err.Error(), sl.Err(err))
		api.Respond(w, r, http.StatusBadRequest, api.ErrResponse{
			Err: err.Error(),
		})

		return
	}

	id, err := s.service.UsersServiceInt.CreateUser(user)
	if err != nil {
		log.Error(err.Error(), sl.Err(err))
		api.Respond(w, r, http.StatusInternalServerError, api.ErrResponse{
			Err: err.Error(),
		})

		return
	}

	log.Info(api.CreationSucceeded)
	api.Respond(w, r, http.StatusOK, api.IdResponse{
		Id:  id,
		Msg: api.CreationSucceeded,
	})

}
