package router

import (
	"net/http"

	"github.com/DimaGlobin/matchme/internal/middleware/auth"
	"github.com/DimaGlobin/matchme/internal/middleware/logger"
	"github.com/DimaGlobin/matchme/internal/server/handler/files_handlers"
	"github.com/DimaGlobin/matchme/internal/server/handler/users_handler"
	"github.com/DimaGlobin/matchme/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
)

func NewRouter(log *slog.Logger, srv *service.Service) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(logger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	//-----------------Authorization---------------------------

	signInHandler := users_handler.NewSignInHandler(log, srv)
	signUpHandler := users_handler.NewSignUpHandler(log, srv)

	router.Post("/sign_up", signUpHandler.Handle())
	router.Post("/sign_in", signInHandler.Handle())

	//---------------------------------------------------------

	//--------------------Private------------------------------

	router.Mount("/api", NewApiRouter(log, srv))

	//--------------------------------------------------------

	return router
}

func NewApiRouter(log *slog.Logger, srv *service.Service) chi.Router {
	router := chi.NewRouter()

	router.Use(auth.New(log, srv))

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	//----------------------Users-----------------------------

	getUserHandler := users_handler.NewGetUserHandler(log, srv)
	updateUserHandler := users_handler.NewUpdateUserHandler(log, srv)
	deleteUserhandler := users_handler.NewDeleteUserHandler(log, srv)

	router.Route("/users", func(r chi.Router) {
		r.Get("/{id}", getUserHandler.Handle())
		r.Get("/", getUserHandler.Handle())
		r.Put("/", updateUserHandler.Handle())
		r.Delete("/", deleteUserhandler.Handle())
	})

	//--------------------------------------------------------

	//-----------------------Photos---------------------------
	uploadFileHandler := files_handlers.NewUploadFileHandler(log, srv)

	router.Route("/photos", func(r chi.Router) {
		r.Post("/", uploadFileHandler.Handle())
	})
	//--------------------------------------------------------
	return router
}
