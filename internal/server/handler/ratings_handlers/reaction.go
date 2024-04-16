package ratings_handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/DimaGlobin/matchme/internal/lib/api"
	"github.com/DimaGlobin/matchme/internal/lib/logger/sl"
	"github.com/DimaGlobin/matchme/internal/middleware/auth"
	"github.com/DimaGlobin/matchme/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
)

type ReactionHandler struct {
	logger  *slog.Logger
	service *service.Service
}

func NewReactionHandler(log *slog.Logger, srv *service.Service) *ReactionHandler {
	return &ReactionHandler{
		logger:  log,
		service: srv,
	}
}

func (rh *ReactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := rh.logger.With(
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	parts := strings.Split(r.URL.Path, "/")
	reaction := parts[len(parts)-2]

	objectIdStr := chi.URLParam(r, "id")
	objectId, err := strconv.ParseUint(objectIdStr, 10, 64)
	if err != nil {
		msg := "Unable to parse id from url query"
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusInternalServerError, msg)

		return
	}

	subjectId := r.Context().Value(auth.UserCtx).(uint64)

	reactionId, matchId, err := rh.service.RatingsServiceInt.AddReaction(reaction, subjectId, objectId)
	if err != nil {
		msg := err.Error()
		log.Error(msg, sl.Err(err))
		api.Respond(w, r, http.StatusInternalServerError, msg)

		return
	}

	log.Info(fmt.Sprintf("Reacted %s successfully, reactionId: %d", reaction, reactionId))
	api.Respond(w, r, http.StatusOK, map[string]interface{}{ //TODO: Bad response body, refactore in future
		reaction + "_id": reactionId,
		"match_id":       matchId,
	})
}
