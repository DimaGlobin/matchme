package ratings_handlers

import (
	"net/http"

	"github.com/DimaGlobin/matchme/internal/lib/api"
	"github.com/DimaGlobin/matchme/internal/lib/logger/sl"
	"github.com/DimaGlobin/matchme/internal/middleware/auth"
	"github.com/DimaGlobin/matchme/internal/service"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
)

type LikeUsersIds struct {
	Likes []uint64 `json:"likes" example:"1,2,3,4"`
}

type GetLikesHandler struct {
	logger  *slog.Logger
	service *service.Service
}

func NewGetLikesHandler(log *slog.Logger, srv *service.Service) *GetLikesHandler {
	return &GetLikesHandler{
		logger:  log,
		service: srv,
	}
}

// @Summary GetLikes
// @Security BearerAuth
// @Tags api
// @Description get all user who liked you
// @ID get-likes
// @Accept  json
// @Produce  json
// @Success 200 {object} LikeUsersIds
// @Failure 500
// @Router /api/action/like/ [get]
func (g *GetLikesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := g.logger.With(
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userId := r.Context().Value(auth.UserCtx).(uint64)

	likes, err := g.service.RatingsServiceInt.GetAllLikes(userId)
	if err != nil {
		msg := "Unable to get likes"
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusInternalServerError, msg)

		return
	}

	log.Info("Likes was successfully sent")
	api.Respond(w, r, http.StatusOK, LikeUsersIds{
		Likes: likes,
	})
}
