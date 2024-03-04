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

func NewSignUpHandler(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &model.User{}

		log := log.With(
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		err := render.DecodeJSON(r.Body, user)
		if err != nil {
			msg := "Unable to decode request body"
			log.Error(msg)
			api.Respond(w, r, http.StatusBadRequest, msg)

			return
		}

		id, err := srv.UsersService.CreateUser(user)
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
}

type SignInBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewSignInHandler(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := &SignInBody{}
		log := log.With(
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		err := render.DecodeJSON(r.Body, body)
		if err != nil {
			msg := "Unable to decode request body"
			log.Error(msg)
			api.Respond(w, r, http.StatusBadRequest, msg)

			return
		}

		token, err := srv.UsersService.GenerateToken(body.Email, body.Password)
		if err != nil {
			msg := "Unable to create jwt token"
			log.Error(err.Error())
			api.Respond(w, r, http.StatusInternalServerError, msg)
		}

		api.Respond(w, r, http.StatusOK, map[string]interface{}{
			"token": token,
		})

	}
}
