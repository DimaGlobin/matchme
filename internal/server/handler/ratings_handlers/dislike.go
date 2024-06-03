package ratings_handlers

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

const (
	dislike = "dislike"
)

type DislikeHandler struct {
	logger  *slog.Logger
	service *service.Service
}

func NewDislikeHandler(log *slog.Logger, srv *service.Service) *DislikeHandler {
	return &DislikeHandler{
		logger:  log,
		service: srv,
	}
}

// @Summary Reaction
// @Security BearerAuth
// @Tags api
// @Description react to user
// @ID dislike
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Dislike
// @Failure 400,401
// @Failure 500
// @Router /api/action/dislike/{id} [post]
func (d *DislikeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := d.logger.With(
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

	subjectId := r.Context().Value(auth.UserCtx).(uint64)

	reactionId, err := d.service.RatingsServiceInt.AddDislike(subjectId, objectId)
	if err != nil {
		msg := err.Error()
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusInternalServerError, msg)

		return
	}

	log.Info(fmt.Sprintf("Disliked successfully, dislikeId: %d", reactionId))
	api.Respond(w, r, http.StatusOK, model.Dislike{
		ReactionType: dislike,
		ReactionId:   reactionId,
	})
}
