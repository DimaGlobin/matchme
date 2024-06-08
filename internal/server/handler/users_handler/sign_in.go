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

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body model.SignInBody true "credentials"
// @Success 200 {object} api.TokenResponse "TokenResponse"
// @Failure 400 {object} api.ErrResponse "ErrResponse"
// @Failure 500 {object} api.ErrResponse "ErrResponse"
// @Router /auth/sign_in [post]
func (s *SigninHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body := &model.SignInBody{}
	log := s.logger.With(
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	err := render.DecodeJSON(r.Body, body)
	if err != nil {
		log.Error(mm_errors.DecodeError.Error(), sl.Err(err))
		api.Respond(w, r, http.StatusBadRequest, api.ErrResponse{
			Err: mm_errors.DecodeError.Error(),
		})

		return
	}

	if err = body.Valid(); err != nil {
		log.Error(err.Error(), sl.Err(err))
		api.Respond(w, r, http.StatusBadRequest, api.ErrResponse{
			Err: err.Error(),
		})

		return
	}

	token, err := s.service.UsersServiceInt.GenerateToken(body.Email, body.Password)
	if err != nil {
		log.Error(mm_errors.JwtCreationError.Error(), sl.Err(err))
		api.Respond(w, r, http.StatusInternalServerError, api.ErrResponse{
			Err: mm_errors.DecodeError.Error(),
		})

		return
	}

	api.Respond(w, r, http.StatusOK, api.TokenResponse{
		Token: token,
	})

}
