package router

import (
	"net/http"

	_ "github.com/DimaGlobin/matchme/docs"
	"github.com/DimaGlobin/matchme/internal/middleware/auth"
	"github.com/DimaGlobin/matchme/internal/middleware/logger"
	"github.com/DimaGlobin/matchme/internal/server/handler/files_handlers"
	"github.com/DimaGlobin/matchme/internal/server/handler/ratings_handlers"
	"github.com/DimaGlobin/matchme/internal/server/handler/users_handler"
	"github.com/DimaGlobin/matchme/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"golang.org/x/exp/slog"
)

func NewRouter(log *slog.Logger, srv *service.Service) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(logger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8084/swagger/doc.json"), //The url pointing to API definition
	))

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	//-----------------Authorization---------------------------

	signInHandler := users_handler.NewSignInHandler(log, srv)
	signUpHandler := users_handler.NewSignUpHandler(log, srv)

	router.Post("/auth/sign_up", signUpHandler.ServeHTTP)
	router.Post("/auth/sign_in", signInHandler.ServeHTTP)

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
	getUserByIdHandler := users_handler.NewGetUserByIdHandler(log, srv)
	updateUserHandler := users_handler.NewUpdateUserHandler(log, srv)
	deleteUserhandler := users_handler.NewDeleteUserHandler(log, srv)

	router.Route("/users", func(r chi.Router) {
		r.Get("/{id}", getUserByIdHandler.ServeHTTP) // TODO: Добавить доп условия на получение пользователей(чтобы любой не мог получить профиль любого)
		r.Get("/", getUserHandler.ServeHTTP)
		r.Put("/", updateUserHandler.ServeHTTP)
		r.Delete("/", deleteUserhandler.ServeHTTP)
	})

	//--------------------------------------------------------

	//-----------------------Photos---------------------------
	uploadFileHandler := files_handlers.NewUploadFileHandler(log, srv)
	getFileByIdHandler := files_handlers.NewGetFileByIdHander(log, srv)
	getFileByNameHandler := files_handlers.NewGetFileByNameHandler(log, srv)
	getFilesHandler := files_handlers.NewGetFilesHandler(log, srv)
	deleteFileHandler := files_handlers.NewDeleteFileHandler(log, srv)

	router.Route("/photos", func(r chi.Router) {
		r.Post("/", uploadFileHandler.ServeHTTP)
		r.Get("/id/{id}", getFileByIdHandler.ServeHTTP)
		r.Get("/", getFilesHandler.ServeHTTP)
		r.Get("/filename/{filename}", getFileByNameHandler.ServeHTTP)
		r.Delete("/{id}", deleteFileHandler.ServeHTTP)
	})

	//--------------------------------------------------------

	//-----------------------Ratings---------------------------

	reactionHandler := ratings_handlers.NewReactionHandler(log, srv)
	rateHadler := ratings_handlers.NewRateUserHandler(log, srv)
	getLikesHandler := ratings_handlers.NewGetLikesHandler(log, srv)

	router.Route("/action", func(r chi.Router) {
		r.Get("/rate", rateHadler.ServeHTTP)
		r.Post("/like/{id}", reactionHandler.ServeHTTP)
		r.Get("/like", getLikesHandler.ServeHTTP)
		r.Post("/dislike/{id}", reactionHandler.ServeHTTP)
	})

	//--------------------------------------------------------
	return router
}
