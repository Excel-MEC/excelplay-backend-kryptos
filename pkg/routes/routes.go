package routes

import (
	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/handlers"
	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/httperrors"
	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/middlewares"
	"github.com/gorilla/mux"
)

// Router wraps mux.Router to add route init method
type Router struct {
	*mux.Router
}

// NewRouter setups and returns a new router
func NewRouter() *Router {
	return &Router{
		mux.NewRouter(),
	}
}

// Routes initializes the routes of the api
func (router *Router) Routes(db *database.DB, config *env.Config) {
	router.HandleFunc("/admin/", handlers.HandleAdmin).Methods("GET")
	router.Handle("/api/ping", middlewares.ErrorsMiddleware(httperrors.Handler(handlers.HeartBeat()))).Methods("GET")
	router.Handle("/api/question",
		middlewares.ErrorsMiddleware(
			httperrors.Handler(
				middlewares.AuthMiddleware(
					handlers.HandleNextQuestion(db, config),
					config,
				),
			),
		),
	).Methods("GET")
	router.Handle("/api/submit",
		middlewares.ErrorsMiddleware(
			httperrors.Handler(
				middlewares.AuthMiddleware(
					handlers.HandleSubmission(db, config),
					config,
				),
			),
		),
	).Methods("POST")
	router.Handle("/api/leaderboard",
		middlewares.ErrorsMiddleware(
			httperrors.Handler(
				middlewares.AuthMiddleware(
					handlers.GetLeaderboard(db, config),
					config,
				),
			),
		),
	).Methods("GET")

	// LoggerMiddleware does not have to be selectively applied because it applies to all endpoints
	router.Use(middlewares.LoggerMiddleware)
}

/*
Note that ErrorsMiddleware must always be the outermost middleware of the selectively applied middlewares
as it handles errors from all internal functions and is the only funcion in the chain that returns the
http.Handler which is expected by router.Handle in it's 2nd arg.
*/
