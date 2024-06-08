package ratings_handlers

import (
	"net/http"

	"github.com/DimaGlobin/matchme/internal/lib/api"
	"github.com/DimaGlobin/matchme/internal/lib/logger/sl"
	"github.com/DimaGlobin/matchme/internal/middleware/auth"
	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/service"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
)

type RateUserHandler struct {
	logger  *slog.Logger
	service *service.Service
}

// @Summary Rate
// @Security BearerAuth
// @Tags api
// @Description get user to rate
// @ID rate
// @Accept  json
// @Produce  json
// @Success 200 {object} model.UserRecommendation
// @Failure 500
// @Router /api/action/rate/ [get]
func NewRateUserHandler(log *slog.Logger, srv *service.Service) *RateUserHandler {
	return &RateUserHandler{
		logger:  log,
		service: srv,
	}
}

func (ru *RateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := ru.logger.With(
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userId := r.Context().Value(auth.UserIdKey).(uint64)

	user, err := ru.service.RatingsServiceInt.RecommendUser(userId)
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
		"validationToken": nil,
	})
}
