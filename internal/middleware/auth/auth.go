package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/DimaGlobin/matchme/internal/lib/api"
	"github.com/DimaGlobin/matchme/internal/service"
	"golang.org/x/exp/slog"
)

const (
	UserCtx = "userId"
)

func New(log *slog.Logger, srv *service.Service) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log := log.With(
			slog.String("component", "middleware/auth"),
		)

		log.Info("auth middleware are enabled")

		fn := func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == "" {
				msg := "Empty auth header"
				api.Respond(w, r, http.StatusUnauthorized, msg)
				log.Error(msg)
				return
			}

			headerParts := strings.Split(header, " ")
			if len(headerParts) != 2 || headerParts[0] != "Bearer" {
				msg := "Invalid auth header"
				api.Respond(w, r, http.StatusUnauthorized, msg)
				log.Error(msg)
				return
			}

			if len(headerParts[1]) == 0 {
				msg := "Token is empty"
				api.Respond(w, r, http.StatusUnauthorized, msg)
				log.Error(msg)
				return
			}

			userId, err := srv.UsersService.ParseToken(headerParts[1])
			if err != nil {
				msg := "Cannot parse token"
				api.Respond(w, r, http.StatusUnauthorized, msg)
				log.Error(msg)
				return
			}

			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), UserCtx, userId)))
		}

		return http.HandlerFunc(fn)
	}
}
