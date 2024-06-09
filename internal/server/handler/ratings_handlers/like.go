package ratings_handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/DimaGlobin/matchme/internal/lib/api"
	"github.com/DimaGlobin/matchme/internal/lib/logger/sl"
	"github.com/DimaGlobin/matchme/internal/middleware/auth"
	"github.com/DimaGlobin/matchme/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
)

type LikeHandler struct {
	logger  *slog.Logger
	service *service.Service
}

func NewLikeHandler(log *slog.Logger, srv *service.Service) *LikeHandler {
	return &LikeHandler{
		logger:  log,
		service: srv,
	}
}

// @Summary Reaction
// @Security BearerAuth
// @Tags api
// @Description react to user
// @ID like
// @Accept  json
// @Produce  json
// @Success 200 {object} model.LikeResp
// @Failure 400,401
// @Failure 500
// @Router /api/action/like/{id} [post]
func (l *LikeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := l.logger.With(
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	objectIdStr := chi.URLParam(r, "id")
	objectId, err := strconv.ParseUint(objectIdStr, 10, 64)
	if err != nil {
		msg := "Unable to parse id from url query"
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusInternalServerError, msg)

		return
	}

	subjectId := r.Context().Value(auth.UserIdKey).(uint64)
	subjectRole := r.Context().Value(auth.UserRoleKey).(string)

	likeResp, err := l.service.RatingsServiceInt.AddLike(subjectId, objectId, subjectRole)
	if err != nil {
		msg := err.Error()
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusInternalServerError, msg)

		return
	}

	log.Info(fmt.Sprintf("Reacted %s successfully, reactionId: %d", likeResp.ReactionType, likeResp.ReactionId))
	api.Respond(w, r, http.StatusOK, likeResp)
}
