package users_handler

import (
	"net/http"
	"strconv"

	"github.com/DimaGlobin/matchme/internal/lib/api"
	"github.com/DimaGlobin/matchme/internal/lib/logger/sl"
	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
)

type GetUserByIdHandler struct {
	logger  *slog.Logger
	service *service.Service
}

// @Summary GetUserById
// @Security BearerAuth
// @Tags api
// @Description get user by id
// @ID get-user-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} model.UserRecommendation
// @Failure 400,401
// @Failure 500
// @Router /api/users/{id} [get]
func NewGetUserByIdHandler(log *slog.Logger, srv *service.Service) *GetUserByIdHandler {
	return &GetUserByIdHandler{
		logger:  log,
		service: srv,
	}
}

func (s *GetUserByIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := s.logger.With(
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	idStr := chi.URLParam(r, "id") // Refactore in future( business logic in habdlers layer )

	var id uint64
	var err error

	id, err = strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		msg := "Unable to parse id from url query"
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusBadRequest, msg)

		return
	}

	user, err := s.service.UsersServiceInt.GetuserById(id)
	if err != nil {
		msg := "Unable to get user"
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusInternalServerError, msg)

		return
	}

	log.Info("User was successfully sent")
	api.Respond(w, r, http.StatusOK, map[string]interface{}{
		"user": model.UserRecommendation{
			Id:          user.Id,
			Name:        user.Name,
			Age:         user.Age,
			Description: user.Description,
		},
	})
}
