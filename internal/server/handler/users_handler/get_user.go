package users_handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/DimaGlobin/matchme/internal/lib/api"
	"github.com/DimaGlobin/matchme/internal/lib/logger/sl"
	"github.com/DimaGlobin/matchme/internal/middleware/auth"
	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/service"
	"github.com/go-chi/chi/v5"
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

func (s *GetUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := s.logger.With(
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	idStr := chi.URLParam(r, "id") // Refactore in future( business logic in habdlers layer )

	fmt.Println("idStr: ", idStr)

	var id uint64
	var err error

	if idStr != "" {
		id, err = strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			msg := "Unable to parse id from url query"
			log.Error(msg, sl.Err(err))
			api.Respond(w, r, http.StatusInternalServerError, msg)

			return
		}
	} else {
		id = r.Context().Value(auth.UserCtx).(uint64)
	}

	user, err := s.service.UsersService.GetuserById(id)
	if err != nil {
		msg := "Unable to get user"
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusInternalServerError, msg)

		return
	}

	if idStr != "" {
		log.Info("User was successfully sent")
		api.Respond(w, r, http.StatusOK, map[string]interface{}{
			"user": model.UserResponse{
				Id:          user.Id,
				Name:        user.Name,
				Age:         user.Age,
				Description: user.Description,
			},
		})

		return
	}

	api.Respond(w, r, http.StatusOK, map[string]interface{}{
		"user": user,
	})
}
