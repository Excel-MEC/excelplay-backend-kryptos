package main

import (
	"database/sql"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	db     *sql.DB
	logger *logrus.Logger
}

func newServer(router *mux.Router, db *sql.DB, logger *logrus.Logger) *server {
	server := &server{
		router,
		db,
		logger,
	}
	server.routes() // Register the route handling functions with the mux router
	return server
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
